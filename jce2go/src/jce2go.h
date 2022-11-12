#ifndef JCE2GO_H
#define JCE2GO_H

#include "parse.h"

#include <cassert>
#include <string>

/**
 * 根据jce生成go文件
 * 包括结构的编解码以及生成Proxy和Servant
 */
class Jce2Go
{
public:
    /**
     * 设置代码生成的根目录
     * @param dir
     */
    void setBaseDir(const string &dir);

    /**
     * 设置包前缀
     * @param prefix
     */
    void setBasePackage(const string &prefix);

    /**
     * 设置gojce 库的路径
     */
    void setGojcePath(const string &gojcePath);

    /**
     * 设置optional 字段默认值是否打包
     */
    void setOptionalPack(bool optional);

    /**
     * 获取optional 字段默认值是否打包
     */
    bool getOptionalPack()const;

    /**
     * 设置tag 属性
     */
    void setTagVec(const vector<string> &attrs);

    /**
     * 获取tags 
     */
    const vector<string>& getTagVec()const;

    /**
     * 生成
     * @param file
     * @param isFramework 是否是框架
     */
    void createFile(const string &file);

protected:
    /**
     * 根据命名空间获取文件路径
     * @param ns 命名空间
     *
     * @return string
     */
    string getFilePath(const string &ns) const;

    string _packagePrefix;
    string _baseDir;
    string _gojcePath;
    bool   _bWithServant;
    bool   _optional;
    vector<string> _tagVec;

    //下面是编解码的源码生成
protected:

    /**
     * 生成某类型的解码源码
     * @param pPtr
     *
     * @return string
     */
    string writeTo(const TypeIdPtr &pPtr) const;

    /**
     * 生成某类型的编码源码
     * @param pPtr
     *
     * @return string
     */
    string readFrom(const TypeIdPtr &pPtr) const;

    /**
     * 
     * @param pPtr
     * 
     * @return string
     */
    string display(const TypeIdPtr &pPtr) const;

    //下面是类型描述的源码生成
protected:

    /*
     * 生成某类型的初始化字符串
     * @param pPtr
     *
     * @return string
     */
    string toTypeInit(const TypePtr &pPtr) const;

    /**
     * 生成某类型的对应对象的字符串描述源码
     * @param pPtr
     *
     * @return string
     */
    string toObjStr(const TypePtr &pPtr, const NamespacePtr &nPtr) const;

    /**
     * 判断是否是对象类型
     */
    bool isObjType(const TypePtr &pPtr) const;

    /**
     * 生成某类型的字符串描述源码
     * @param pPtr
     *
     * @return string
     */
    string tostr(const TypePtr &pPtr, const NamespacePtr &nPtr) const;

    /**
     * 生成内建类型的字符串源码
     * @param pPtr
     *
     * @return string
     */
    string tostrBuiltin(const BuiltinPtr &pPtr) const;
    /**
     * 生成vector的字符串描述
     * @param pPtr
     *
     * @return string
     */
    string tostrVector(const VectorPtr &pPtr, const NamespacePtr &nPtr) const;

    /**
     * 生成map的字符串描述
     * @param pPtr
     *
     * @return string
     */
    string tostrMap(const MapPtr &pPtr, const NamespacePtr &nPtr) const;

    /**
     * 生成某种结构的符串描述
     * @param pPtr
     *
     * @return string
     */
    string tostrStruct(const StructPtr &pPtr, const NamespacePtr &nPtr) const;

    /**
     * 生成某种枚举的符串描述
     * @param pPtr
     *
     * @return string
     */
    string tostrEnum(const EnumPtr &pPtr) const;

    /**
     * 生成类型变量的解码源码
     * @param pPtr
     *
     * @return string
     */
    string decode(const TypeIdPtr &pPtr) const;

    /**
     * 生成类型变量的编码源码
     * @param pPtr
     *
     * @return string
     */
    string encode(const TypeIdPtr &pPtr) const;

    /**
     * 转换id为go源码形势
     * @param pPtr
     *
     * @return string
     */
    string parseMemberId(const string &strId) const;

    /**
     * 获取成员变量的package
     * @param pPtr
     *
     * @return string
     */
    string getMemberNamespace(const TypePtr &pPtr) const;
    
    /**
     * 生成首字母大写的函数名
     * @param name
     *
     * @return string
     */
    string initialUpper(const string &name) const;

    //以下是h和go文件的具体生成
protected:
    /**
     * 生成结构的go文件内容
     * @param pPtr
     *
     * @return string
     */
    string generateGo(const StructPtr &pPtr, const NamespacePtr &nPtr) const;

    /**
     * 生成容器的go源码
     * @param pPtr
     *
     * @return string
     */
    string generateGo(const ContainerPtr &pPtr) const;

    /**
     * 生成枚举的头文件源码
     * @param pPtr
     *
     * @return string
     */
    string generateGo(const EnumPtr &pPtr, const NamespacePtr &nPtr) const;

    /**
     * 生成常量go源码
     * @param pPtr
     * 
     * @return string
     */
    void generateGo(const ConstPtr &pPtr, const NamespacePtr &nPtr) const;

    void generateGo(const vector<EnumPtr> &es,const vector<ConstPtr> &cs,const NamespacePtr &nPtr) const;
    /**
     * 生成名字空间go文件源码
     * @param pPtr
     *
     * @return string
     */
    void generateGo(const NamespacePtr &pPtr) const;

    /**
     * 生成每个jce文件的go文件源码
     * @param pPtr
     *
     * @return string
     */
    void generateGo(const ContextPtr &pPtr) const;
	/**
	*  生成interface的go文件源码
	*/
    string generateGo(const InterfacePtr &pPtr , const NamespacePtr &nPtr) const;
    
	string generateGo(const OperationPtr& pPtr, const InterfacePtr &iPtr,const NamespacePtr &nPtr) const;
    
	string generateGo(const OperationPtr& pPtr, const NamespacePtr &nPtr) const;
    
	string generateGoCase(const OperationPtr& pPtr, const NamespacePtr &nPtr) const;
    /**
     * 生成参数声明的go文件内容
     * @param pPtr
     *
     * @return string
     */
    string generateGo(const ParamDeclPtr &pPtr, const NamespacePtr &nPtr) const;

};

#endif // JCE2GO_H

