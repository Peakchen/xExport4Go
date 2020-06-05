#encoding: utf-8 

# add by stefan 20191109 14:33 for develop.
# xls export go file.

import xlrd 
from collections import OrderedDict
import json
import codecs
import glob
import os,sys
import os.path
import re

import sys
reload(sys)
sys.setdefaultencoding("utf8")

from define import ExportType

totaltabName = "total"
totaltabs = {}

# purpose: read file and start export
# param1: need export aim file.
# param2: filter condition.
def loopExportXld(dstformat, filter):
    print os.getcwd()
    for file in glob.glob("*.xls"):
        wb = xlrd.open_workbook(file)
        sheets = wb.sheets()
        #print "show sheets list: "
        for sheet in sheets:
            exportSheet(sheet, dstformat, filter)

def exportSheet(sheet, dstformat, filter):
    #print sheet.name
    shn = sheet.name
    shn = shn.encode("UTF-8")  
    if shn == totaltabName:
        exportTotalTab(sheet)
    else:
        print shn
        export2File(sheet, dstformat, filter)

# purpose: export total table from others exported.
# param: tab sheet.

def exportTotalTab(sheet):
    print sheet.name
    rows = sheet.get_rows()
    rowidx = 0
    for row in rows:
        rowidx = rowidx + 1
        if rowidx == 1:
            continue
        
        tablename = row[0].value
        if tablename == '':
            continue

        isexport = row[1].value
        begincol = int(row[4].value)
        endcol = int(row[5].value)
        totaltabs[tablename] = {'isexport':isexport, 'begincol': begincol, 'endcol': endcol} 
        #print totaltabs

# begin export operation.
# param1: tab sheet.
# param2: need export aim file.
# param3: filter condition.
def export2File(sheet, dstformat, filter):
    convertlist = []
    shn = sheet.name
    shn = shn.encode("UTF-8")  
    if totaltabs.has_key(shn) == False:
        return

    expitem = totaltabs[shn]
    exportflag = expitem['isexport']
    if exportflag.lower() == "no":
        return

    #print expitem
    rows = sheet.get_rows()
    rowidx = 0
    rowitemtypes = OrderedDict()
    colitemtypes = OrderedDict()
    totallimit = expitem['endcol']
    for row in rows:
        rowidx = rowidx + 1
        if row[0].value == '':
            continue

        if row[0].value == 'CLOSE':
            break

        if rowidx == 2 and row[1].value != shn: # if not exist, then exist.
            return

        if rowidx > 7:
            break

        colidx = 0
        for col in row:
            colidx = colidx + 1
            colval = col.value
            if colval == '' and colidx == 2:
                continue
        
            strRowidx = str(rowidx)
            strcolidx = str(colidx)

            if totallimit < colidx:
                continue

            if rowitemtypes.has_key('3') == True and rowidx >= 7:
                rowdata = rowitemtypes['3']
                if rowdata.has_key(strcolidx) == True:
                    if len(rowdata) == totallimit and rowdata[strcolidx] == filter:
                        continue
            
            if isinstance(colval, str):
                colval = colval.encode("UTF-8") 

            if rowitemtypes.has_key(strRowidx) == False:
                if rowidx >= 7 and colidx > 1:
                    fldname = rowitemtypes['6']
                    fldtype = rowitemtypes['4']
                    val = {strcolidx: [{"FLDNAME": fldname[strcolidx]}, {"FLDTYPE": fldtype[strcolidx]}]}
                    if totaltabs.has_key(strRowidx) == False:
                        rowitemtypes[strRowidx] = val
                    else:
                        rowitemtypes[strRowidx].update(val)  

                    if colitemtypes.has_key(strcolidx) == False:
                        colitemtypes[strcolidx] = val

                elif rowidx < 7:
                    rowitemtypes[strRowidx] = {strcolidx: colval}
            else:
                if rowidx >= 7 and colidx > 1:
                    fldname = rowitemtypes['6']
                    fldtype = rowitemtypes['4']
                    val = {strcolidx: [{"FLDNAME": fldname[strcolidx]}, {"FLDTYPE": fldtype[strcolidx]}]}
                    rowitemtypes[strRowidx].update(val)  
                    if colitemtypes.has_key(strcolidx) == False:
                        colitemtypes[strcolidx] = val
                        colitemtypes.pop(strcolidx)

                elif rowidx < 7:
                    rowitemtypes[strRowidx].update({strcolidx: colval})  


    convertlist = colitemtypes['3']
    createRow = rowitemtypes['2'] # get table file, judge it can be export.
    if totaltabs.has_key(createRow['2']) == True and len(convertlist) > 0:
        print "find it, it can be export."
        producefile(shn,shn+dstformat, convertlist)
    

def producefile(struct, filename, rowlist):
    print filename
    gofile = "package Config\n\n\n"
    JsonFile = struct + ".json"
    gofile += "/* \n" +  "\t\texport from " + JsonFile + " by tool.\n*/\n"  
    struct = str(struct)
    fileStruct = "T" + struct.capitalize() + "Base"
    gofile += "type\t" + fileStruct + "\t struct {\n"
    rowidx = 0
    # build structure: type TXXX struct { ... }
    for idx in rowlist:
        rowidx = rowidx + 1
        item = rowlist[str(idx)]
        fldname = item[0]['FLDNAME']
        fldtype = item[1]['FLDTYPE']
        fldname = str(fldname)
        gofile += "\t" + fldname.capitalize() + "\t" + ExportType[fldtype] + "`json:" + '"' + fldname + '"' + '`'
        if rowidx == len(rowlist):
            gofile += "\n}\n\n\n"
        else:
            gofile += "\n"

    #build struct func:
    configStruct = "T" + struct.capitalize()
    gofile += "type\t" + configStruct + "\t struct {  data []* "+ fileStruct + " \n}\n\n"
    exUseConf = "\tG" + struct.capitalize()
    arrStructConf = "tArr" + struct.capitalize() 
    gofile += "type\t" + arrStructConf + "\t" + "[]*" + fileStruct + "\n\n"
    gofile += "var (\n" + exUseConf + "\t*" + configStruct + "=" + "&" + configStruct + "{}\n)\n\n"
    gofile += "func init(){\n}\n\n"
    gofile += "func load"+ struct.capitalize() +"() {\n\tvar(\n\t\tpath string\n\t)\n\tif len(SvrPath) == 0 {\n\t\tpath = getserverpath()\n\t} \n\tpath = filepath.Join(SvrPath, "+ '"' + JsonFile + '"' +") \n\t" + "Config.ParseJson2Cache(" + exUseConf + "," + "&" + arrStructConf + "{}" + ", path)\n}\n\n"
    gofile += "func (this *" + configStruct + ") ComfireAct(data interface{}) (errlist []string) { errlist = []string{} \n\tcfg := data.(*" + arrStructConf + ") \n\tfor _, item := range *cfg {\n\n\t}\n\treturn\n}\n\n"
    gofile += "func (this *" + configStruct + ") DataRWAct(data interface{}) (errlist []string) {\n\tcfg := data.(*" + arrStructConf + ") \n\tthis.data = []*"+fileStruct+"{}\n\tfor _, item := range *cfg {\n\n\t}\n\treturn\n}\n\n"
    gofile += "func (this *" + configStruct + ") Get(idx int) *"+fileStruct+"{\n\tif idx >= len(this.data) {\n\t\treturn nil\n\t}\n\treturn this.data[idx]\n}"
    #build new file for go code.
    with codecs.open(filename, "w", "utf-8") as target:
        target.write(gofile)

def _main_():
    loopExportXld(".go", 's')

_main_()
