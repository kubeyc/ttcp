#ifndef TTCP_INET_ADDRESS_H
#define TTCP_INET_ADDRESS_H

#include "common.h"

class InetAddress: copyable
{
public:
    InetAddress(const StringArg& ip, in_port_t port);

    InetAddress(in_port_t port, bool loopback);

    InetAddress(const sockaddr_in& addr);

    std::string toIP() const;

    std::string toEndpoint() const;

    inline const sockaddr_in& getSockAddr() const { return __sockaddr; };

    inline void setSockAddr(const sockaddr& addr) { std::memcpy(&__sockaddr, &addr, sizeof addr); };
private:
    sockaddr_in __sockaddr;
};


#endif /* TTCP_INET_ADDRESS_H */