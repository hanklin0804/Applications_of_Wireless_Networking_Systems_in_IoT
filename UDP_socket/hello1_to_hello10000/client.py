import socket
import time
localIP     = "140.118.122.155"
localPort   = 5407



# Create a datagram socket
UDPClientSocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)

# Bind to address and ip
UDPClientSocket.bind((localIP, localPort))

print("UDP Client up and listening")

time_start = time.time()
# Listen for incoming datagrams
n = 0
while(n< 10000):
    # Sending a msg to Proxy
    n += 1
    msg = 'hello '+ str(n)
    UDPClientSocket.sendto(msg.encode(), ("140.118.122.155", 5406))


    clientMsg, clientIP = UDPClientSocket.recvfrom(1024)

    print("Message from Sever: ",clientMsg.decode())
    print("Sever IP Address:", clientIP,"\n")

print(time.time()-time_start)

