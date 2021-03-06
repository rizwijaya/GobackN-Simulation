# Sender.py
import time, socket, sys

def decimalToBinary(n):  
    return n.replace("0b", "")

def binary(s):
    a_byte_array = bytearray(s, "utf8")

    byte_list = []

    for byte in a_byte_array:
        binary_representation = bin(byte)
        byte_list.append(decimalToBinary(binary_representation))

    #print(byte_list)
    a=""
    for i in byte_list:
        a=a+i
    return a

def connecting():
    print("\nWelcome to Chat Room\n")
    print("Initialising....\n")
    time.sleep(1)

    s = socket.socket()
    host = socket.gethostname()
    ip = socket.gethostbyname(host)
    port = 1234
    s.bind((host, port))
    print(host, "(", ip, ")\n")
    name = input(str("Enter your name: "))
            
    s.listen(1)
    print("\nWaiting for incoming connections...\n")
    conn, addr = s.accept()
    print("Received connection from ", addr[0], "(", addr[1], ")\n")

    s_name = conn.recv(1024)
    s_name = s_name.decode()
    print(s_name, "has connected to the chat room\nEnter [e] to exit chat room\n")
    conn.send(name.encode())
    return conn

def main():
    conn = connecting()
    while True:
        message = input(str("Me : "))
        conn.send(message.encode())
        if message == "[e]":
            message = "Left chat room!"
            conn.send(message.encode())
            print("\n")
            break
        message=binary(message)
        f=str(len(message))
        conn.send(f.encode())
    
        i=0
        j=0
        print("Panjang pesan: " + f)
        j=int(input("Enter the window size -> "))
    
        b=""
    
        j=j-1
        f=int(f)
        k=j
        while i!=f:
            while(i!=(f-j)):
                conn.send(message[i].encode())
                b=conn.recv(1024)
                b=b.decode()
                print(b)
                if(b!="ACK Lost"):
                    time.sleep(1)
                    print("Acknowledgement Diterima! Sliding window pada rentang "+(str(i+1))+" ke "+str(k+1)+" Sekarang kirim paket selanjutnya")
                    i=i+1
                    k=k+1
                    time.sleep(1)
                else:
                    time.sleep(1)
                    print("Acknowledgement data bit nya LOST! Sliding window pada rentang "+(str(i+1))+" ke "+str(k+1)+" Sekarang kirim ulang paket yang sama")
                    time.sleep(1)
            while(i!=f):
                
                conn.send(message[i].encode())
                b=conn.recv(1024)
                b=b.decode()
                print(b)
                if(b!="ACK Lost"):
                    time.sleep(1)
                    print("Acknowledgement Diterima! Sliding window pada rentang "+(str(i+1))+" ke "+str(k)+" Sekarang kirim paket selanjutnya")
                    i=i+1
                    time.sleep(1)
                else:
                    time.sleep(1)
                    print("Acknowledgement data bit nya LOST! Sliding window pada rentang "+(str(i+1))+" ke "+str(k)+" Sekarang kirim ulang paket yang sama")
                    time.sleep(1)

if __name__ == "__main__":
    main()            
     
