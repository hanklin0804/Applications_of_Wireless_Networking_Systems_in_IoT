import socket

localIP     = "127.0.0.1"
localPort   = 5405



# Create a datagram socket
UDPSeverSocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)

# Bind to address and ip
UDPSeverSocket.bind((localIP, localPort))

print("UDP Sever up and listening")

# Listen for incoming datagrams
while(True):
    try:
        SeverMsg, SeverIP = UDPSeverSocket.recvfrom(1024)
        print("Message from Proxy: ",SeverMsg.decode())
        print("Proxy IP Address: ",SeverIP,"\n")

        msg = 'World ' +str(SeverMsg.decode().split(' ')[1])
        
        # Sending a msg to Client
        UDPSeverSocket.sendto(msg.encode(), ("127.0.0.1", 5407))
    except:
        # Create a datagram socket
        UDPSeverSocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)

        # Bind to address and ip
        UDPSeverSocket.bind((localIP, localPort))