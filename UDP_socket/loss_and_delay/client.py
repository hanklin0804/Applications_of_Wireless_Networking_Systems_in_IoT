import socket
import time
import select
localIP     = "140.118.122.155"
localPort   = 5407



# Create a datagram socket
UDPClientSocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)
# Bind to address and ip
UDPClientSocket.bind((localIP, localPort))

print("UDP Client up and listening")


# Listen for incoming datagrams
n = 0
clientMsg = ''
clientIP = ''
start_time = time.time()
while(True):
    if n < 10000:
        # Sending a msg to Proxy
        n += 1
        msg = 'hello '+ str(n)
        UDPClientSocket.sendto(msg.encode(), ("140.118.122.155", 5406))

    
    # timeout setting 10ms
    ready = select.select([UDPClientSocket], [], [], 0.01)
    if ready[0]:
        clientMsg, clientIP = UDPClientSocket.recvfrom(1024)

    print("Message from Sever: ",clientMsg.decode())
    print("Sever IP Address:", clientIP,"\n")
    if int(clientMsg.decode().split(' ')[1]) >= 10000:
        break
print( time.time() - start_time )





