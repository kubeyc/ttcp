#ifndef TTCP_COMMON_H
#define TTCP_COMMON_H

#include <utility>
#include <cstring>
#include <cassert>
#include <string>

class nocopyable
{
protected:
    nocopyable() = default;
    ~nocopyable() = default;
private:
    nocopyable(const nocopyable&) = delete;
    nocopyable& operator=(const nocopyable&) = delete;
};

struct copyable
{

};

class StringArg: copyable
{
public:
    StringArg(const char* arg) {
        __arg = new char[ std::strlen(arg) + 1];
        std::strcpy(__arg, arg);
    }

    ~StringArg() {
        delete []__arg;
    }

    StringArg(const StringArg& that) {
        __arg = new char[ std::strlen(that.__arg) + 1];
        std::strcpy(__arg, that.__arg);
    }

    StringArg& operator=(const StringArg& that) {
        if (this != &that) {
            delete []__arg;
            __arg = new char[ std::strlen(that.__arg) + 1];
            std::strcpy(__arg, that.__arg);
        } 

        return *this;
    }

    StringArg(StringArg&& that) {
        __arg = that.__arg;
    }

    StringArg& operator=(StringArg&& that) {
        if (&that != this) {
            __arg = that.__arg;
        }
        
        return *this;
    }

    inline const char* get_c_str() const { return __arg; };

private:
    char* __arg;
};

#endif /*TTCP_COMMON_H*/