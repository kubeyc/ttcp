#ifndef TTCP_TCP_STREAM_H
#define TTCP_TCP_STREAM_H

#include "common.h"
#include "socket.h"

class TcpStream :nocopyable
{
public:
    explicit TcpStream(Socket&& socket);

    int receiveAll();

    int sendAll();

private:
    Socket __sock;
};


#endif /*TTCP_TCP_STREAM_H*/