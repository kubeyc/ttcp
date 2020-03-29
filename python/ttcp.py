import socket
import sys
import struct
import time

def recv_all(client_socket, n, bufsize):
    readn = 0
    total_data = []
    while readn < n:
        data = client_socket.recv(bufsize)
        if len(data) == 0:
            return 0
        
        if len(data) > 0:
            readn += len(data)
            total_data += data
    
    return len(total_data)

def send_all(sock, data, n):
    sendn = 0
    while sendn < n:
        nw = sock.sendall(struct.pack("%ds" % len(data), data))
        if nw == 0:
            return 0
        
        if nw > 0:
            sendn += nw

    return sendn

def transmit(server_socket, number, length):
    payloadMsgData = bytearray()
    for i in range (length):
        payloadMsgData.append(b"0123456789ABCDEF"[i%16])

    server_socket.send(struct.pack('=2I', socket.htonl(number), socket.htonl(length)))
    
    for i in range(number):
        server_socket.send(struct.pack('=I', socket.htonl(length)))
        
        server_socket.sendall(struct.pack("s", payloadMsgData))
        # if nw == 0:
        #     print("send_all error, server closed")
        #     server_socket.close()
        #     return
        
        # if nw != length:
        #     print("send_all %d data, need %d data" % (length, length))
        #     server_socket.close()
        #     return

        ack_data = server_socket.recv(4)
        if len(ack_data) == 0:
            print("recv error, server closed")
            server_socket.close()
            return

        ack = socket.ntohl(struct.unpack('=I', ack_data)[0])
        if ack != length:
            print("send_all %d data, server recv %d data" % (nw, ack))
            server_socket.close()
            return


    print("send ok")
    server_socket.close()

def receive(server_socket, client_socket):
    data = client_socket.recv(8)
    number = socket.ntohl(struct.unpack('=I', data[:4])[0])
    length = socket.ntohl(struct.unpack('=I', data[4:8])[0])

    print("receive buffer length = %d\nreceive number of buffers = %d" % (length, number))

    totalMib = number * length * 1.0 / 1024 / 1024
    print("%.3f MiB in total" % (totalMib))

    start_time = time.time()

    for i in range (number):
        payloadLengthData = client_socket.recv(4)
        if len(payloadLengthData) == 0:
            print("client closed")
            client_socket.close()
            server_socket.close()
            return

        payloadMsgLength = socket.ntohl(struct.unpack('=I', payloadLengthData[:4])[0])
        nr = recv_all(client_socket, payloadMsgLength, 8192)
        if nr == 0:
            client_socket.close()
            server_socket.close()
            return
        
        if nr != payloadMsgLength:
            print("receive payload message length: %d, need length: %d" % (len(payloadMsgData), payloadMsgLength))
            client_socket.close()
            server_socket.close()
            return
        
       
        
        nw = client_socket.send(struct.pack('=I', socket.htonl(payloadMsgLength)))
        print(nw)


    end_time = time.time()
    elapsed = end_time - start_time
    print("%.3f seconds\n%.3f MiB/s\n" % (elapsed, totalMib / elapsed))
    client_socket.close()
    server_socket.close()

def main(argv):
    if len(argv) < 3:
        print ("Usage:\n -t (client or server)")
        return

    if argv[1] != "-t":
        print ("Usage:\n -t (client or server)")
        return

    if argv[2] == "client":
        pass

    if argv[2] == "server":
        host = ""
        port = 0
        for i in range(len(argv)):
            if argv[i] == "-h" and i + 1 < len(argv):
                host = argv[i+1]
            if argv[i] == "-p" and i + 1 < len(argv):
                port = int(argv[i+1])

        if host == "":
            print ("Usage:\n -t server -h host")
            return

        if port == 0:
            print ("Usage:\n -t server -h %s -p port" % (host))
            return

        print ("ttcp listen server %s:%d" % (host, port))
        
        server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        server_socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        server_socket.setsockopt(socket.SOL_SOCKET, socket.TCP_NODELAY, 1)
        server_socket.bind((host, port))
        server_socket.listen(5)
        (client_socket, client_address) = server_socket.accept()
        receive(server_socket, client_socket)
  
    if argv[2] == "client":
        host = ""
        port = 0
        number = 1000
        length = 1024
        for i in range(len(argv)):
            if argv[i] == "-h" and i + 1 < len(argv):
                host = argv[i+1]
            if argv[i] == "-p" and i + 1 < len(argv):
                port = int(argv[i+1])
            if argv[i] == "-n" and i + 1 < len(argv):
                number = int(argv[i+1])
            if argv[i] == "-b" and i + 1 < len(argv):
                length = int(argv[i+1])

        if host == "":
            print ("Usage:\n -t client -h host")
            return

        if port == 0:
            print ("Usage:\n -t client -h %s -p port" % (host))
            return
        
        print ("ttcp connection server %s:%d:\n  message number: %d count\n  message length: %d b\n  total: %.3f Mib" % (host, port, number, length, 
            float(1.0 * number * length / 1024 / 1024)) )

        server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        server_socket.setsockopt(socket.SOL_SOCKET, socket.TCP_NODELAY, 1)
        server_socket.connect((host, port))
        transmit(server_socket, number, length)

if __name__ == "__main__":
    main(sys.argv)