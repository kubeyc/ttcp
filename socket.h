#ifndef TTCP_SOCKET_H
#define TTCP_SOCKET_H

#include "common.h"
#include "inet_address.h"
#include <sys/socket.h>

class Socket :nocopyable
{

public:
    explicit Socket(int sockfd): __sockfd(sockfd) {};
    
    Socket(Socket&& that); // socket 移动拷贝构造

    Socket& operator=(Socket&& that); // Socket 移动赋值重载构造

    ~Socket();

    inline const int get_sockfd() const { return __sockfd; };

    void bind(const InetAddress& addr);
    
    void listen(int backlog);

    void setRefuse(bool refuse);

    static Socket createTCP();    

private:
    int __sockfd;
};


#endif /* TTCP_SOCKET_H */