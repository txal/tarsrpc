#include "util/tc_option.h"
#include "util/tc_common.h"

namespace taf
{

void TC_Option::decode(int argc, char *argv[])
{
    _mParam.clear();

    vector<string> v;
    for(int i = 1; i < argc; i++)
    {
        v.push_back(argv[i]);
    }

    for(size_t i = 0; i < v.size(); i++)
    {
        if(v[i].length() > 2 && v[i].substr(0,2) == "--")
        {
            parse(v[i]);
        }
        else
        {
            _vSingle.push_back(v[i]);
        }
    }
}

void TC_Option::parse(const string &s)
{
    string::size_type pos = s.find('=');
    if( pos != string::npos)
    {
        _mParam.insert(STORE_TYPE::value_type(s.substr(2, pos-2), s.substr(pos+1)));
    }
    else
    {
        _mParam.insert(STORE_TYPE::value_type(s.substr(2, pos-2), ""));
    }
}

string TC_Option::getValue(const string &sName)
{
    STORE_TYPE::const_iterator it = _mParam.find(sName);
    if(it != _mParam.end())
    {
        return it->second;
    }
    return "";
}

vector<string> TC_Option::getValues(const string &sName)
{
    vector<string> values;
    std::pair<STORE_TYPE::const_iterator, STORE_TYPE::const_iterator> it = _mParam.equal_range(sName);
    for (; it.first != it.second; ++it.first) {
        values.push_back(it.first->second);
    }
    return values;
}

bool TC_Option::hasParam(const string &sName)
{
    return _mParam.find(sName) != _mParam.end();
}

vector<string>& TC_Option::getSingle()
{
    return _vSingle;
}

TC_Option::STORE_TYPE& TC_Option::getMulti()
{
    return _mParam;
}

}

