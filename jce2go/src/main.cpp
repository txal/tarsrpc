#include "util/tc_option.h"
#include "util/tc_file.h"
#include "jce2go.h"

static string gojce_path = "jce_parser/gojce";

void usage()
{
    cout << "Usage : jce2go [OPTION] jcefile" << endl;
    cout << "  jce2go support type: bool byte short int long float double vector map" << endl;
    cout << "supported [OPTION]:" << endl;
    cout << "  --help                help,print this(帮助)" << endl;
    cout << "  --dir=DIRECTORY       generate golang file to DIRECTORY(生成文件到目录DIRECTORY,默认为当前目录)" << endl;
    cout << "  --gopath=DIRECTORY    generate golang package DIRECTORY(生成包的GOPATH相对路径,默认为空)" << endl;
    cout << "  --no-optional (不打包option字段)" << endl;
    cout << "  --gojce-path (jce-parser路径 默认为" << gojce_path << ")" << endl;
    cout << "  --add-tag (设置默认struct tag)" << endl;
    cout << endl;
    exit(0);
}

void check(vector<string> &vJce)
{
    for(size_t i  = 0; i < vJce.size(); i++)
    {
        string ext  = taf::TC_File::extractFileExt(vJce[i]);
        if(ext == "jce")
        {
            if(!taf::TC_File::isFileExist(vJce[i]))
            {
                cerr << "file '" << vJce[i] << "' not exists" << endl;
                usage();
                exit(0);
            }
        }
        else
        {
            cerr << "only support jce file." << endl;
            exit(0);
        }
    }
}

int main(int argc, char* argv[])
{
    if(argc < 2)
    {
        usage();
    }

    taf::TC_Option option;
    option.decode(argc, argv);
    vector<string> vJce = option.getSingle();

    check(vJce);

    if(option.hasParam("help"))
    {
        usage();
    }

    Jce2Go j2go;


    //设置生成文件的根目录
    if(option.getValue("dir") != "")
    {
        j2go.setBaseDir(option.getValue("dir"));
    }
    else
    {
        j2go.setBaseDir(".");
    }

    //包名前缀
    if(option.hasParam("base-package"))
    {
        j2go.setBasePackage(option.getValue("base-package"));
    }
    else
    {
        j2go.setBasePackage("");
    }

    //是否打包 optional 字段
    j2go.setOptionalPack(true);
    if(option.hasParam("no-optional"))
    {
        j2go.setOptionalPack(false);
    }

    // 设置gojce 库的路径
    if(option.hasParam("gojce-path"))
    {
        string path = option.getValue("gojce-path");
        if (!path.empty()) gojce_path = path;
    }
    j2go.setGojcePath(gojce_path);

    // 设置结构tag
    if(option.hasParam("add-tag"))
    {
        j2go.setTagVec(option.getValues("add-tag"));
    }

    try
    {
        //是否可以以taf开头
        g_parse->setTaf(option.hasParam("with-taf"));
        g_parse->setHeader(option.getValue("gopath"));

        for(size_t i = 0; i < vJce.size(); i++)
        {
            g_parse->parse(vJce[i]);
            j2go.createFile(vJce[i]);
        }
    }catch(exception& e)
    {
        cerr<<e.what()<<endl;
    }

    return 0;
}
