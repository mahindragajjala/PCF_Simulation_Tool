https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.91d2rkfy9hlw

To capture the trace of a mobile device with a laptop using Wireshark, follow this step-by-step guide:

1. Prerequisites
Laptop Setup: Install Wireshark on your laptop. Download it from https://www.wireshark.org/.
Network Setup: Ensure the mobile and laptop are connected to the same network (Wi-Fi or mobile hotspot).
Administrative Privileges: Run Wireshark as an administrator to access network interfaces.
Wireless Adapter: If you're capturing Wi-Fi traffic, ensure your laptop's wireless adapter supports monitor mode (required to capture unencrypted packets).

2. Flow for Capturing the Trace
A. Network Topology
Connection Type:


Option 1: Use your mobile's Wi-Fi to connect to a shared network with your laptop.
Option 2: Use your mobile hotspot and connect your laptop to it.
IP Addressing: Note the mobile's IP address. This is crucial for filtering the traffic later.


You can find the mobile's IP address in the Wi-Fi settings (e.g., 192.168.x.x).

B. Setting Up Wireshark
Start Wireshark:


Open Wireshark and select the correct network interface (e.g., Wi-Fi adapter or Ethernet if you're using a tethering cable).
Set Display Filters:


Use a filter like:
 ip.addr == <mobile's IP address>
 Replace <mobile's IP address> with your mobile's actual IP.
Enable Monitor Mode (Optional):


If capturing all network traffic (e.g., in public Wi-Fi), enable monitor mode on your wireless adapter to capture packets from all devices.

C. Capturing Mobile Activity During a Call
Power Cycle the Mobile:


Switch off the mobile and restart it. This ensures the device reinitializes its network connection and generates new traffic.
Start the Capture:


Begin capturing packets in Wireshark before switching on the mobile.
Make the Call:


After the mobile connects to the network, place a call to another device. This action generates traffic between the mobile and the network.
Stop the Capture:


Once the call ends, stop the capture in Wireshark.

3. Analyzing the Trace
A. Filter Traffic
Apply filters to isolate relevant data:
To focus on specific protocols: sip, rtp, http, etc.
To isolate mobile IP traffic:
 ip.addr == <mobile's IP>


B. Look for Key Events
DHCP Requests: Look for DHCP packets when the mobile obtains an IP address.
Filter: dhcp
DNS Queries: Check DNS queries for domain name resolution.
Filter: dns
VoIP Traffic: Analyze SIP (Session Initiation Protocol) or RTP (Real-Time Protocol) packets if it's a VoIP call.
Filter: sip or rtp
C. Export Logs
Export the trace for further analysis:
Go to File > Export Packet Dissections > As Plain Text or save the .pcap file for sharing or deeper analysis.

4. Important Notes
Encryption:


Traffic from apps or services (e.g., WhatsApp, Telegram) is usually encrypted, so packet contents may not be visible unless decrypted keys are available.
VoLTE traffic may also require specific network tools to decode LTE traffic.
Legal Compliance:


Ensure you have legal permission to capture traffic, especially if monitoring devices that are not yours.
Advanced Tools:


If the network involves LTE or 5G, you may need additional tools like QXDM, tshark, or specialized hardware.

Let me know if you'd like examples of specific Wireshark filters or additional guidance for your setup!

