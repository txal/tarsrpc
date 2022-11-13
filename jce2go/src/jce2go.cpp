#include "jce2go.h"
#include "util/tc_md5.h"
#include "util/tc_file.h"
#include "util/tc_common.h"
#include <string>
#include <string.h>
#include <set>

#define JCE_PACKAGE     "gojce"

#define TAB g_parse->getTab()
#define INC_TAB g_parse->incTab()
#define DEL_TAB g_parse->delTab()

//////////////////////////////////////////////////////////////////////////////////
// string Jce2Go::toTypeInit(const TypePtr &pPtr) const
// {
//     BuiltinPtr bPtr = BuiltinPtr::dynamicCast(pPtr);
//     if (bPtr)
//     {
//         switch (bPtr->kind())
//         {
//             case Builtin::KindBool:     return "false";
//             case Builtin::KindByte:     return "0";
//             case Builtin::KindShort:    return "0";
//             case Builtin::KindInt:      return "0";
//             case Builtin::KindLong:     return "0";
//             case Builtin::KindFloat:    return "0.0f";
//             case Builtin::KindDouble:   return "0.0";
//             case Builtin::KindString:   return "\"\";";
//             default:                    return "";
//         }
//     }

//     VectorPtr vPtr = VectorPtr::dynamicCast(pPtr);
//     if (vPtr)
//     {
//         //数组特殊处理
//         string sType;
//         size_t iPosBegin, iPosEnd;
//         sType = tostr(vPtr->getTypePtr());
//         if ((iPosBegin = sType.find("<")) != string::npos && (iPosEnd = sType.rfind(">")) != string::npos)
//         {
//             sType = sType.substr(0, iPosBegin) +  sType.substr(iPosEnd+1);
//         }
//         //[] (数组)的数组变为[1]
//         sType = taf::TC_Common::replace(sType, "[]" , "[1]");
//         return "new([]" + sType + ", 1)";;
//     }

//     MapPtr mPtr = MapPtr::dynamicCast(pPtr);
//     if (mPtr) return "make(" + tostrMap(mPtr, true) + ")";

//     StructPtr sPtr = StructPtr::dynamicCast(pPtr);
//     if (sPtr) return "new(" + tostrStruct(sPtr) + ")";

//     EnumPtr ePtr = EnumPtr::dynamicCast(pPtr);
//     if (ePtr) return "0";

//     return "";
// }

string Jce2Go::toObjStr(const TypePtr &pPtr, const NamespacePtr &nPtr) const
{
    string sType = tostr(pPtr, nPtr);

    // if (sType == "bool") return "bool";
    // if (sType == "byte")    return "byte";
    // if (sType == "short" )  return "int16";
    // if (sType == "int" )    return "int32";
    // if (sType == "long" )   return "int64";
    // if (sType == "float" )  return "float32";
    // if (sType == "double" ) return "float64";

    return sType;
}

string Jce2Go::tostr(const TypePtr &pPtr, const NamespacePtr &nPtr) const
{
    BuiltinPtr bPtr = BuiltinPtr::dynamicCast(pPtr);
    if (bPtr) return tostrBuiltin(bPtr);

    VectorPtr vPtr = VectorPtr::dynamicCast(pPtr);
    if (vPtr) return tostrVector(vPtr, nPtr);

    MapPtr mPtr = MapPtr::dynamicCast(pPtr);
    if (mPtr) return tostrMap(mPtr, nPtr);

    StructPtr sPtr = StructPtr::dynamicCast(pPtr);
    if (sPtr) return tostrStruct(sPtr, nPtr);

    EnumPtr ePtr = EnumPtr::dynamicCast(pPtr);
    if (ePtr) return tostrEnum(ePtr);

    if (!pPtr) return "interface{}";

    assert(false);
    return "";
}

/*******************************BuiltinPtr********************************/
string Jce2Go::tostrBuiltin(const BuiltinPtr &pPtr) const
{
    string s;

    switch (pPtr->kind())
    {
        case Builtin::KindBool:     s = "bool";     break;
        case Builtin::KindByte:     s = "byte";     break;
        case Builtin::KindShort:    s = "int16";    break;
        case Builtin::KindInt:      s = "int32";    break;
        case Builtin::KindLong:     s = "int64";    break;
        case Builtin::KindFloat:    s = "float32";  break;
        case Builtin::KindDouble:   s = "float64";  break;
        case Builtin::KindString:   s = "string";   break;
        case Builtin::KindVector:   s = "[]";       break;
        case Builtin::KindMap:      s = "map";      break;
        default:                    assert(false);  break;
    }

    return s;
}

string Jce2Go::tostrVector(const VectorPtr &pPtr, const NamespacePtr &nPtr) const
{
    string s = string("[]") + toObjStr(pPtr->getTypePtr(), nPtr);
    return s;
}

string Jce2Go::tostrMap(const MapPtr &pPtr, const NamespacePtr &nPtr) const
{
    string s;
    s = string("map[") + toObjStr(pPtr->getLeftTypePtr(), nPtr) + "]" + toObjStr(pPtr->getRightTypePtr(), nPtr);

    return s;
}

string Jce2Go::tostrStruct(const StructPtr &pPtr, const NamespacePtr &nPtr) const
{
    //return taf::TC_Common::replace(pPtr->getSid(), "::", ".");
    string prefix = pPtr->getSid().substr(0, pPtr->getSid().find("::"));
    if (prefix == nPtr->getId())
    {
        return taf::TC_Common::replace(pPtr->getSid(), "::", ".").substr(taf::TC_Common::replace(pPtr->getSid(), "::", ".").find(".")+1);
    } 
    else 
    {
        return taf::TC_Common::replace(pPtr->getSid(), "::", ".");
    }
}

string Jce2Go::tostrEnum(const EnumPtr &pPtr) const
{
    return "int";
}

string Jce2Go::parseMemberId(const string &strId) const
{
    //return "M_" + strId;
    return initialUpper(strId);
}

string Jce2Go::getMemberNamespace(const TypePtr &pPtr) const
{
    VectorPtr vPtr = VectorPtr::dynamicCast(pPtr);
    if (vPtr) return getMemberNamespace(vPtr->getTypePtr());

    MapPtr mPtr = MapPtr::dynamicCast(pPtr);
    if (mPtr) return getMemberNamespace(mPtr->getRightTypePtr());

    StructPtr sPtr = StructPtr::dynamicCast(pPtr);
    if (sPtr)
    {
        return sPtr->getSid().substr(0, sPtr->getSid().find("::"));
    }
    return "";
}
#include <iostream>

/******************************StructPtr***************************************/
string Jce2Go::generateGo(const StructPtr &pPtr, const NamespacePtr &nPtr) const
{
    ostringstream s;
    s << g_parse->printHeaderRemark();

    vector<string> key = pPtr->getKey();
    vector<TypeIdPtr>& member = pPtr->getAllMemberPtr();

    s << TAB << "package "<< nPtr->getId()<< endl;
    s << TAB << "import \"reflect\"" << endl;
    s << TAB << _gojcePath << endl;

    set<string> importNames;
    for (size_t i = 0; i < member.size(); i++)
    {
        string name = getMemberNamespace(member[i]->getTypePtr());
        if (name != "" && name != nPtr->getId() && importNames.find(name) == importNames.end())
        {
            importNames.insert(name);
            s << TAB << "import \"" << g_parse->getHeader() << name <<"\"" << endl;
        }
    }
    s << endl;

    s << TAB << "type " << pPtr->getId() << " struct {"<< endl;
    INC_TAB;

    //定义成员变量set;get函数
    for (size_t i = 0; i < member.size(); i++)
    {
        string sId = member[i]->getId();
        string sTag = member[i]->getTagAttr();
        const vector<string> &tags = getTagVec();
        if (sTag.empty() && !tags.empty()) {
            sTag = " `";
            for (vector<string>::const_iterator itb = tags.begin(), ite = tags.end();
                    itb != ite; ++itb) {
                sTag.append(*itb).append(":\"");
                sTag.append(sId);
                sTag.append("\" ");
            }
            sTag.append("`");
        }
        s << TAB << parseMemberId(sId) << " "<< tostr(member[i]->getTypePtr(), nPtr) << sTag << endl;
        //s << TAB << "`json:\"" << member[i]->getId() << "\"`";
        //s << endl;
    }

    DEL_TAB;
    s << TAB << "}" << endl;
    s << endl;

    //resetDefault()
    s << TAB << "func (_obj *"<< pPtr->getId() << ") resetDefault() {" << endl;
    INC_TAB;
    for (size_t i = 0; i < member.size(); i++)
    {
        if (member[i]->hasDefault())
        {
            BuiltinPtr bPtr  = BuiltinPtr::dynamicCast(member[i]->getTypePtr());

            if (bPtr && bPtr->kind() == Builtin::KindString)
            {
                s << TAB << "_obj." << parseMemberId(member[i]->getId()) << " = \"" << 
                    taf::TC_Common::replace(member[i]->def(), "\"", "\\\"") <<"\"" << endl;
            }
            else if (bPtr && bPtr->kind() == Builtin::KindVector)
            {
                s << TAB << "_obj." << parseMemberId(member[i]->getId()) << " = make([]" << 
                    member[i]->getTypePtr() << ", 0)" << endl;
            }
            else if (bPtr && bPtr->kind() == Builtin::KindMap)
            {
                MapPtr mPtr = MapPtr::dynamicCast(member[i]->getTypePtr());
                if (mPtr)
                {
                    s << TAB << "_obj." << parseMemberId(member[i]->getId()) << " = make(map[" << 
                        toObjStr(mPtr->getLeftTypePtr(), nPtr) << "]" << toObjStr(mPtr->getRightTypePtr(), nPtr) << ")" << endl;
                }
            }
            else
            {
                s << TAB << "_obj." << parseMemberId(member[i]->getId()) << " = " << member[i]->def() << endl;
            }
        }
    }
    DEL_TAB;
    s << TAB << "}" << endl;
    s << endl;

    //WriteTo()
    s << TAB << "func (_obj *"<< pPtr->getId() << ") WriteTo(_os "<< JCE_PACKAGE << ".JceOutputStream) error {" << endl;
    INC_TAB;
    s << TAB << "var _err error" << endl;
    bool optional = getOptionalPack();
    for (size_t i = 0; i < member.size(); i++)
    {
        TypePtr tPtr = member[i]->getTypePtr();
        BuiltinPtr bPtr  = BuiltinPtr::dynamicCast(tPtr);
        VectorPtr vPtr  = VectorPtr::dynamicCast(tPtr);
        MapPtr mPtr  = MapPtr::dynamicCast(tPtr);
        bool require = member[i]->isRequire() or optional or !(bPtr or vPtr or mPtr);
        if (!require) {
            if (vPtr or mPtr) {
                s << TAB << "if len(_obj." << parseMemberId(member[i]->getId())
                    << ") != 0";
            } else switch (bPtr->kind()) {
                case Builtin::KindVoid:
                    continue;
                case Builtin::KindBool:
                case Builtin::KindByte:
                case Builtin::KindShort:
                case Builtin::KindInt:
                case Builtin::KindLong:
                case Builtin::KindFloat:
                case Builtin::KindDouble:
                    s << TAB << "if _obj." << parseMemberId(member[i]->getId())
                        << " != " << member[i]->def();
                    break;
                case Builtin::KindString:
                    s << TAB << "if _obj." << parseMemberId(member[i]->getId())
                        << " != \"" <<
                    taf::TC_Common::replace(member[i]->def(), "\"", "\\\"") << "\"";
                    break;
                case Builtin::KindVector:
                case Builtin::KindMap:
                    s << TAB << "if len(_obj." << parseMemberId(member[i]->getId())
                        << ") != 0";
                    break;
            }
            s << " {" << endl;
            INC_TAB;
        }
        s << TAB << "if _err = _os.Write(reflect.ValueOf(&_obj." << parseMemberId(member[i]->getId()) << "), "
        << member[i]->getTag() << "); _err != nil {" << endl;
        INC_TAB;
        s << TAB << "return _err" << endl;
        DEL_TAB;
        s << TAB << "}" << endl;
        if (!require) {
            DEL_TAB;
            s << TAB << "}" << endl;
        }
    }
    s << TAB << "return nil" << endl;
    DEL_TAB;
    s << TAB << "}" << endl;
    s << endl;

    //ReadFrom()
    s << TAB << "func (_obj *"<< pPtr->getId() << ") ReadFrom(_is "<< JCE_PACKAGE << ".JceInputStream) error {" << endl;
    INC_TAB;
    s << TAB << "var _err error" << endl;
    s << TAB << "var _i interface{}" << endl;
    s << TAB << "_obj.resetDefault()" << endl;
    for (size_t i = 0; i < member.size(); i++)
    {
        s << TAB << "_i, _err = _is.Read(reflect.TypeOf(_obj." << parseMemberId(member[i]->getId()) << "), " << member[i]->getTag() << ", " 
            << (member[i]->isRequire() ? "true" : "false") << ")" << endl;
        s << TAB << "if _err != nil {" << endl;
        INC_TAB;
        s << TAB << "return _err" << endl;
        DEL_TAB;
        s << TAB << "}" << endl;
        s << TAB << "if _i != nil {" << endl;
        INC_TAB;
        s << TAB << "_obj." << parseMemberId(member[i]->getId()) << " = _i.(" << tostr(member[i]->getTypePtr(), nPtr) << ")" << endl;
        DEL_TAB;
        s << TAB << "}" << endl;
    }
    s << TAB << "return nil" << endl;
    DEL_TAB;
    s << TAB << "}" << endl;
    s << endl;

    //Display()
    s << TAB << "func (_obj *"<< pPtr->getId() << ") Display(_ds "<< JCE_PACKAGE << ".JceDisplayer) {" << endl;
    INC_TAB;
    for (size_t i = 0; i < member.size(); i++)
    {
        s << TAB << "_ds.Display(reflect.ValueOf(&_obj." << parseMemberId(member[i]->getId()) << "), \"" << member[i]->getId() << "\")" << endl;
    }
    DEL_TAB;
    s << TAB << "}" << endl;
    s << endl;

    //WriteJson
    s << TAB << "func (_obj *"<< pPtr->getId() << ") WriteJson(_en "<< JCE_PACKAGE << ".JceJsonEncoder) ([]byte, error) {" << endl;
    INC_TAB;
    s << TAB << "var _err error" << endl;
    for (size_t i = 0; i < member.size(); i++)
    {
        s << TAB << "_err = _en.EncodeJSON(reflect.ValueOf(&_obj." << parseMemberId(member[i]->getId()) << "), \"" << member[i]->getId() << "\")" << endl;
        s << TAB << "if _err != nil {" << endl;
        INC_TAB;
        s << TAB << "return nil, _err" << endl;
        DEL_TAB;
        s << TAB << "}" << endl;
    }
    s << TAB << "return _en.ToBytes(), nil" << endl;
    DEL_TAB;
    s << TAB << "}" << endl;
    s << endl;

    //ReadJson
    s << TAB << "func (_obj *"<< pPtr->getId() << ") ReadJson(_de "<< JCE_PACKAGE << ".JceJsonDecoder) error {" << endl;
    INC_TAB;
    s << TAB << "return _de.DecodeJSON(reflect.ValueOf(_obj))" << endl;
    DEL_TAB;
    s << TAB << "}" << endl;
    s << endl;

    string fileCs  = getFilePath(nPtr->getId()) + pPtr->getId() + ".go";
    taf::TC_File::makeDirRecursive(getFilePath(nPtr->getId()), 0755);
    taf::TC_File::save2file(fileCs, s.str());

    return s.str();
}

/******************************ConstPtr***************************************/
void Jce2Go::generateGo(const vector<EnumPtr> &es,const vector<ConstPtr> &cs,const NamespacePtr &nPtr) const
{
    if (es.size()==0 && cs.size()==0)
    {
        return;
    }
    ostringstream s;
    s << g_parse->printHeaderRemark();

    s << TAB << "package " << nPtr->getId()<< endl;

    if (cs.size()>0)
    {
        s << TAB << "const (" << endl;
        INC_TAB;
        //-----------------const类型开始------------------------------------
        for (size_t i = 0; i < cs.size(); i++)
        {
            if (cs[i]->getConstTokPtr()->t == ConstTok::STRING)
            {
                string tmp = taf::TC_Common::replace(cs[i]->getConstTokPtr()->v, "\"", "\\\"");
                s  << TAB << cs[i]->getTypeIdPtr()->getId()<< " = \"" << tmp << "\""<< endl;
            }
            else
            {
                s  << TAB << cs[i]->getTypeIdPtr()->getId()<< " = " << cs[i]->getConstTokPtr()->v << endl;
            }
        }
        DEL_TAB;
        s << TAB << ")" << endl;
    }
    //-----------------const类型结束--------------------------------
    if (es.size()>0)
    {
        //-----------------枚举类型开始---------------------------------
        for (size_t i = 0; i < es.size(); i++)
        {
            s << TAB << "const ("<<endl;
            INC_TAB;
            vector<TypeIdPtr>& member = es[i]->getAllMemberPtr();
            for (size_t i = 0; i < member.size(); i++)
            {
                s << TAB << member[i]->getId();
                if(member[i]->hasDefault())
                {
                    s << " = " << member[i]->def() << endl;
                }
                else
                {
                    s << " = iota" << endl;
                }
            }
            DEL_TAB;
            s<< TAB <<")"<<endl;
        }
    }
    //-----------------枚举类型结束---------------------------------

    string fileCs  = getFilePath(nPtr->getId()) + nPtr->getId()+"_const.go";
    taf::TC_File::makeDirRecursive(getFilePath(nPtr->getId()), 0755);
    taf::TC_File::save2file(fileCs, s.str());

    return;
}

/******************************NamespacePtr***************************************/
void Jce2Go::generateGo(const NamespacePtr &pPtr) const
{
    vector<StructPtr>       &ss    = pPtr->getAllStructPtr();
    vector<EnumPtr>         &es    = pPtr->getAllEnumPtr();
    vector<ConstPtr>        &cs    = pPtr->getAllConstPtr();
    //string                  &na    = pPtr->getId();

    //interface
    vector<InterfacePtr>    &is    = pPtr->getAllInterfacePtr();
    for (size_t i = 0; i < is.size(); i++)
    {
        generateGo(is[i], pPtr);
    }

    for (size_t i = 0; i < ss.size(); i++)
    {
        generateGo(ss[i], pPtr);
    }

    generateGo(es,cs,pPtr);//go里面的枚举、const都放到一起。

    return;
}

void Jce2Go::generateGo(const ContextPtr &pPtr) const
{
    vector<NamespacePtr> namespaces = pPtr->getNamespaces();

    vector<string> include = pPtr->getIncludes();
    for (size_t i = 0; i < include.size(); i++)
    {
        //TODO 依赖解析，比较复杂，需要解析文件，并取出module name。
        //还要判断所有的参数，确定哪个interface或者结构体引用了依赖，否则go无法编译
        //cout << taf::TC_Common::replace(taf::TC_File::extractFileName(include[i]), ".h","") << endl;
		cout << "!!!do not support include now " <<endl;
        cout << include[i] << endl;
    }
    for (size_t i = 0; i < namespaces.size(); i++)
    {
        if(namespaces[i]->hasInterface())
        {
            //TODO 不知道做啥好
        }
    }

    for (size_t i = 0; i < namespaces.size(); i++)
    {
        generateGo(namespaces[i]);
    }
}

void Jce2Go::createFile(const string &file)
{
    std::vector<ContextPtr> contexts = g_parse->getContexts();
    for (size_t i = 0; i < contexts.size(); i++)
    {
        if (file == contexts[i]->getFileName())
        {
            generateGo(contexts[i]);
        }
    }
}
/*************************InterfacePtr*************************/

struct SortOperation {
    bool operator ()(const OperationPtr& o1, const OperationPtr& o2)
    {
        return o1->getId() < o2->getId();
    }
};

string Jce2Go::generateGo(const InterfacePtr &pPtr , const NamespacePtr &nPtr) const
{
    ostringstream s;
    vector<OperationPtr>& vOperation = pPtr->getAllOperationPtr();
    std::sort(vOperation.begin(), vOperation.end(), SortOperation());
    // upper name
    string name = initialUpper(pPtr->getId());
    s << g_parse->printHeaderRemark();
    s << TAB << "package " << nPtr->getId()<< endl;
    s << TAB << "import (" << endl;
    INC_TAB;
    //s << TAB << "\"fmt\"" << endl;
    s << TAB << "\"code.com/tars/goframework/jce/taf\"" << endl;
    s << TAB << "\"code.com/tars/goframework/jce_parser/gojce\"" << endl;
    s << TAB << "\"reflect\"" << endl;
    //s << TAB << ". \"taf/servant\"" << endl;
    s << TAB << "m \"code.com/tars/goframework/tars/servant/model\""<< endl;
    s << TAB << "\"errors\""<< endl;
    s << TAB << "context \"context\""<<endl;
    s << TAB << ")" << endl;
    s << "type "<<name<<" struct {"<<endl;
    s << TAB <<"s m.Servant"<<endl;
    s << "}"<<endl;
    DEL_TAB;

    //proxy
    for (size_t i = 0; i < vOperation.size(); i++)
    {
        s << generateGo(vOperation[i],pPtr, nPtr) << endl;
    }
    
    INC_TAB;
    s << "func(_obj *"<<name<<") SetServant(s m.Servant){"<<endl;
    s <<TAB<<"_obj.s = s"<<endl;
    s <<"}"<<endl;
    DEL_TAB;

    //server
    //
    //type interface
    s << "type _imp"<<name<<" interface {"<<endl;
    for (size_t i = 0; i < vOperation.size(); i++)
    {
        s << generateGo(vOperation[i], nPtr);
    }
    s << "}"<<endl;

    //dispatch
    INC_TAB;
    s << "func(_obj *"<<name<<") Dispatch(ctx context.Context,_val interface{}"<<", ";
    s << "req *taf.RequestPacket) (*taf.ResponsePacket,error){" <<endl;
    //如果jce接口全都没有入参，parms参数将不会有地方用到会导致编译报错
    bool useparms=false;
    for (size_t i = 0; i < vOperation.size(); i++)
    {
        vector<ParamDeclPtr>& vParamDecl = vOperation[i]->getAllParamDeclPtr();
        if (vParamDecl.size()>0)
        {
            useparms=true;
            break;
        }
    }
    if (useparms){
        s <<TAB<<"parms := gojce.NewInputStream(req.SBuffer)"<<endl;
    }
    s <<TAB<<"oe := gojce.NewOutputStream()"<<endl;
    s <<TAB<<"_imp := _val.(_imp"<<name<<")"<<endl;
    s <<TAB<<"switch req.SFuncName {"<<endl;
    for (size_t i = 0; i < vOperation.size(); i++)
    {
        s << generateGoCase(vOperation[i], nPtr);
    }
    INC_TAB;
    s <<TAB<<"default:"<< endl;
    INC_TAB;
    s <<TAB<<"return nil,errors.New(\"func mismatch\")"<<endl;
    DEL_TAB;
    DEL_TAB;
    s <<TAB<< "}"<<endl;
    s <<TAB<<"var status map[string]string"<<endl;
    s <<TAB<< "return &taf.ResponsePacket{" <<endl;
    INC_TAB;
    s <<TAB<< "IVersion:     1,"<<endl;
    s <<TAB<< "CPacketType:  0,"<<endl;
    s <<TAB<< "IRequestId:   req.IRequestId,"<<endl;
    s <<TAB<< "IMessageType: 0,"<<endl;
    s <<TAB<< "IRet:         0,"<<endl;
    s <<TAB<< "SBuffer:      oe.ToBytes(),"<<endl;
    s <<TAB<< "Status:       status,"<<endl;
    s <<TAB<< "SResultDesc:  \"\","<<endl;
    s <<TAB<< "Context:      req.Context,"<<endl;
    DEL_TAB;
    s <<TAB<<"},nil"<<endl;
    //s <<TAB<<"return nil"<<endl;
    s << "}"<<endl;
    DEL_TAB;

    //对interface接口加一个单独的文件名后缀_IF，方便查询，反正golang不认文件名
    string fileCs  = getFilePath(nPtr->getId()) + pPtr->getId() + "_IF.go";
    taf::TC_File::makeDirRecursive(getFilePath(nPtr->getId()), 0755);
    taf::TC_File::save2file(fileCs, s.str());
    return s.str();
}
string Jce2Go::generateGoCase(const OperationPtr &pPtr, const NamespacePtr &nPtr) const
{
    ostringstream s;
    vector<ParamDeclPtr>& vParamDecl = pPtr->getAllParamDeclPtr();
    INC_TAB;
    s <<TAB<< "case \""<<pPtr->getId()<<"\":"<<endl;
    INC_TAB;
    for (size_t i = 0; i < vParamDecl.size(); i++)
    {
        if (!vParamDecl[i]->isOut())
        {
            s <<TAB<<"var p_"<<i<<" "<<tostr(vParamDecl[i]->getTypeIdPtr()->getTypePtr(), nPtr)<<endl;
            //t_1, err := parms.Read(reflect.TypeOf(p_1), 1, true)
            s <<TAB<<"t_"<<i<<",err := parms.Read(reflect.TypeOf(p_"<<i<<"),"<<i+1<<",true)"<<endl;
            s <<TAB<<"if err != nil{"<<endl;
            INC_TAB;
            //s <<TAB<<"fmt.Println(err.Error())"<<endl;
            s <<TAB<<"return nil,err"<<endl;
            DEL_TAB;
            s <<TAB<<"}"<<endl;
        }
        else
        {
            s <<TAB<<"var o_"<<i<<" "<<tostr(vParamDecl[i]->getTypeIdPtr()->getTypePtr(), nPtr)<<endl;
        }
    }
    //return  ret
    if (pPtr->getReturnPtr()->getTypePtr())
    {
        //s <<TAB<<"var _ret " << tostr(pPtr->getReturnPtr()->getTypePtr(), nPtr)<< endl;
        s <<TAB<<"_ret,err := _imp."<<initialUpper(pPtr->getId())<<"(";
    }
    else
    {
        if(vParamDecl.size()>0){
            // err already used
            s << TAB<<"err = _imp."<<initialUpper(pPtr->getId())<<"(";
        }else{
            s << TAB<<"err := _imp."<<initialUpper(pPtr->getId())<<"(";
        }

    }
    for (size_t i = 0; i < vParamDecl.size(); i++)
    {
        if (!vParamDecl[i]->isOut())
        {
            s <<"t_"<<i<<".("<<tostr(vParamDecl[i]->getTypeIdPtr()->getTypePtr(), nPtr)<<")";
        }
        else
        {
            s <<"&o_"<<i;
        }
        if ((i+1)<vParamDecl.size()){
             s <<",";
        }
    }
    s <<")"<<endl;

    s <<TAB<<"if err != nil{"<<endl;
    INC_TAB;
    s <<TAB<<"return nil,err"<<endl;
    DEL_TAB;
    s <<TAB<<"}"<<endl;
    //encode ret
    if (pPtr->getReturnPtr()->getTypePtr())
    {
        s <<TAB<<"oe.Write(reflect.ValueOf(&_ret), 0)"<<endl;
    }
    for (size_t i = 0; i < vParamDecl.size(); i++)
    {
        if (vParamDecl[i]->isOut())
        {
            s <<TAB<<"oe.Write(reflect.ValueOf(&o_"<<i<<"),"<<i+1<<")"<<endl;
        }
    }
    DEL_TAB;
    DEL_TAB;
    return s.str();
}
string Jce2Go::generateGo(const OperationPtr &pPtr, const NamespacePtr &nPtr) const
{
    ostringstream s;
    vector<ParamDeclPtr>& vParamDecl = pPtr->getAllParamDeclPtr();
    INC_TAB;
    s << TAB <<initialUpper(pPtr->getId());
    s << "(";
    for (size_t i = 0; i < vParamDecl.size(); i++)
    {
        if ((i+1) == vParamDecl.size()){
            s << generateGo(vParamDecl[i],nPtr);
        }else{
            s << generateGo(vParamDecl[i],nPtr) << ",";
        }
    }
    s << ")";
    s << " (";
    if (pPtr->getReturnPtr()->getTypePtr())
    {
        s << tostr(pPtr->getReturnPtr()->getTypePtr(), nPtr)<<",";
    }
    s << "error)";
    s << endl;
    DEL_TAB;
    return s.str();
}

string Jce2Go::generateGo(const OperationPtr &pPtr,const InterfacePtr &iPtr ,const NamespacePtr &nPtr) const
{
    ostringstream s;
    vector<ParamDeclPtr>& vParamDecl = pPtr->getAllParamDeclPtr();
    INC_TAB;
    // func 
    s <<"func (_obj *" << initialUpper(iPtr->getId())<<") ";
    s <<initialUpper(pPtr->getId());
    s <<"(";
    string routekey = "";
    for (size_t i = 0; i < vParamDecl.size(); i++)
    {
        s << generateGo(vParamDecl[i], nPtr) << ",";

        if (routekey.empty() && vParamDecl[i]->isRouteKey())
        {
            routekey = vParamDecl[i]->getTypeIdPtr()->getId();
        }

    }
    s << "_opt ...map[string]string )";
    if (pPtr->getReturnPtr()->getTypePtr())
    {
        s <<"(_ret " << tostr(pPtr->getReturnPtr()->getTypePtr(), nPtr)<< ",_err error)";
    }
    else
    {
        // if no return,return error
        s <<"error";
    }
    s <<"{"<<endl;
    //func content
    s <<TAB<< "_oe := gojce.NewOutputStream()"<<endl;
    for (size_t i = 0; i < vParamDecl.size(); i++)
    {
        if (!vParamDecl[i]->isOut())
        {
            s << TAB << "_oe.Write(reflect.ValueOf(&";
            s <<vParamDecl[i]->getTypeIdPtr()->getId() <<"), ";
            s << i+1;
            s << ")"<<endl;
        }
    }
    s << TAB << "var (" << endl;
    INC_TAB;
    if (pPtr->getReturnPtr()->getTypePtr()){
        s << TAB << "_resp *taf.ResponsePacket" << endl;
    }
    s << TAB << "err error" << endl;
//    s << TAB << "_status map[string]string"<<endl;
//    s << TAB << "_context map[string]string"<<endl;
    DEL_TAB;
    s << TAB << ")"<<endl;
    if (pPtr->getReturnPtr()->getTypePtr()){
        s << TAB << "_resp,err = _obj.s.Taf_invoke(context.TODO(),0,\""<<pPtr->getId()<<"\", _oe.ToBytes())"<<endl;
    }else{
        s << TAB << "_,err = _obj.s.Taf_invoke(context.TODO(),0,\""<<pPtr->getId()<<"\", _oe.ToBytes())"<<endl;
    }
    s << TAB << "if err != nil {"<<endl;
    INC_TAB;
    //s <<TAB<< "fmt.Println(err.Error()) "<<endl;
    if (pPtr->getReturnPtr()->getTypePtr())
    {
        s << TAB << "return _ret,err"<<endl;
    }
    else
    {
        s << TAB << "return err"<<endl;
    }
    DEL_TAB;
        s <<TAB<<"}"<<endl;
    //func resp
    //如果没有参数和返回值，定义了is编译就报错
    bool haveOut=false;
    for (size_t i = 0; i < vParamDecl.size(); i++)
    {
        if (vParamDecl[i]->isOut()){
            haveOut = true;
            break;
        }
    }
    if (haveOut || pPtr->getReturnPtr()->getTypePtr()){
        s << TAB << "_is := gojce.NewInputStream(_resp.SBuffer)"<<endl;
    }
    //end
    if (pPtr->getReturnPtr()->getTypePtr()){
        s << TAB << "r0, err := _is.Read(reflect.TypeOf(_ret), 0 ,true)"<<endl;
        s <<TAB<< "if err!=nil {"<<endl;
        INC_TAB;
        //s <<TAB<< "fmt.Println(err.Error()) "<<endl;
        if (pPtr->getReturnPtr()->getTypePtr())
        {
            s << TAB << "return _ret,err"<<endl;
        }
        else
        {
            s << TAB << "return err"<<endl;
        }
        DEL_TAB;
        s <<TAB<<"}"<<endl;
    }
    for (size_t i = 0; i < vParamDecl.size(); i++)
    {
        if (vParamDecl[i]->isOut())
        {
            //r3, err := is.Read(reflect.TypeOf(*c), 3 ,true)
            s << TAB <<"r_"<<i+1 <<", err := _is.Read(reflect.TypeOf(*";
            s <<vParamDecl[i]->getTypeIdPtr()->getId();
            s << "),"<<i+1<<",true)"<<endl;
            s <<TAB<< "if err!=nil {"<<endl;
            INC_TAB;
            //s <<TAB<< "fmt.Println(err.Error()) "<<endl;
            if (pPtr->getReturnPtr()->getTypePtr())
            {
                s << TAB << "return _ret,err"<<endl;
            }
            else
            {
                s << TAB << "return err"<<endl;
            }
            DEL_TAB;
            s <<TAB<<"}"<<endl;
            //convert interface
            s <<TAB<<"tmp_"<<i+1<<":="<<"r_"<<i+1;
            s <<".("<<tostr(vParamDecl[i]->getTypeIdPtr()->getTypePtr(), nPtr)<<")"<<endl;
            s <<TAB<<"*"<<vParamDecl[i]->getTypeIdPtr()->getId()<<"="<<"tmp_"<<i+1<<endl;
        }
    }
    if (pPtr->getReturnPtr()->getTypePtr())
        s <<TAB <<"return r0.("<<tostr(pPtr->getReturnPtr()->getTypePtr(), nPtr)<<"),nil"<<endl;
    else
        s <<TAB <<"return nil"<<endl;
    s <<"}";
    DEL_TAB;
    return s.str();
}

string Jce2Go::generateGo(const ParamDeclPtr& pPtr, const NamespacePtr &nPtr) const
{
    ostringstream s;

    s << pPtr->getTypeIdPtr()->getId()<<" ";
    //if (pPtr->isOut() || pPtr->getTypeIdPtr()->getTypePtr()->isSimple())
    if (pPtr->isOut())
    {
        s << "*";
        s << tostr(pPtr->getTypeIdPtr()->getTypePtr(), nPtr);
    }
    else
    {
    //结构, map, vector, string
        s << tostr(pPtr->getTypeIdPtr()->getTypePtr(), nPtr) ;
    }
    return s.str();
}

/************************end InterfacePtr***********************************/

void Jce2Go::setBaseDir(const string &dir)
{
    _baseDir = dir;
}

void Jce2Go::setBasePackage(const string &prefix)
{
    _packagePrefix = prefix;
    if (_packagePrefix.length() != 0 && _packagePrefix.substr(_packagePrefix.length()-1, 1) != ".")
    {
        _packagePrefix += ".";
    }
}

void Jce2Go::setGojcePath(const string &gojcePath)
{
    _gojcePath = "import \"";
    _gojcePath.append("code.com/tars/goframework/").append(gojcePath).append("\"");
}

void Jce2Go::setOptionalPack(bool optional)
{
    _optional = optional;
}

bool Jce2Go::getOptionalPack()const
{
    return _optional;
}

void Jce2Go::setTagVec(const vector<string> &attrs)
{
    _tagVec = attrs;
}

const vector<string>& Jce2Go::getTagVec()const
{
    return _tagVec;
}

string Jce2Go::getFilePath(const string &ns) const
{
    return _baseDir + "/" + taf::TC_Common::replace(_packagePrefix, ".", "/") + "/" + ns + "/";
}

string Jce2Go::initialUpper(const string &name) const
{
    string u = name;
    if ((u[0] >= 'a') && (u[0] <= 'z'))  
        u[0] = u[0] + ('A' - 'a');  
    return u;
}
