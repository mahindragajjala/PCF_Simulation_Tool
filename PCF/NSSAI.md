https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.lx0mtdbps0uw




In 5G networks, NSSAI (Network Slice Selection Assistance Information) is an essential concept for enabling network slicing, which allows the creation of multiple virtual networks (or slices) on the same physical network infrastructure. Each slice can be tailored to meet the specific requirements of different applications, such as enhanced mobile broadband (eMBB), ultra-reliable low-latency communication (URLLC), or massive machine-type communication (mMTC).
Here’s a breakdown of NSSAI components and an example:

NSSAI Components
NSSAI consists of one or more Single NSSAI (S-NSSAI) entries. Each S-NSSAI is defined by:
Slice/Service Type (SST):


Identifies the type of network slice.
Examples:
SST = 1: eMBB
SST = 2: URLLC
SST = 3: mMTC
Slice Differentiator (SD) (Optional):


Provides further differentiation within the same SST.
A 24-bit value (e.g., 0x112233).

Example of NSSAI
Let’s consider a scenario where a 5G network provides slices for different services:
eMBB Slice:


SST = 1 (eMBB for high-speed data, e.g., video streaming)
SD = 0x010203 (specific to a particular operator or region)
URLLC Slice:


SST = 2 (low-latency service for applications like autonomous driving)
SD = 0x112233 (specific to a factory automation use case)
mMTC Slice:


SST = 3 (massive IoT for smart metering and sensors)
SD = Not provided (default slice for general mMTC services)

Practical Example
NSSAI for a 5G Device
A 5G device (UE) may include the following NSSAI in its communication with the network:
{
  "defaultNSSAI": [
    {
      "sst": 1,
      "sd": "0x010203"
    },
    {
      "sst": 2,
      "sd": "0x112233"
    },
    {
      "sst": 3
    }
  ]
}

This NSSAI tells the network that the device requires:
A slice for eMBB services (e.g., video streaming),
A slice for URLLC services (e.g., remote surgery or autonomous driving),
A slice for mMTC services (e.g., IoT sensors).

How NSSAI Works in Practice
UE Registration:


The UE sends its NSSAI to the network during the registration process.
Example: A user in a smart factory might send an NSSAI indicating the need for URLLC slices for factory automation.
AMF (Access and Mobility Management Function):


The AMF determines the appropriate network slices based on the NSSAI.
If the requested slices are available, the network allocates resources.
Slice-Specific Resource Allocation:


Once slices are selected, the UE is granted access to the specific slices, ensuring tailored QoS (Quality of Service) for its applications.

Code Example in Go (Mock Example)
Here’s a simple Go example to represent NSSAI handling in a 5G system:
package main

import (
	"fmt"
)

// Define NSSAI structure
type NSSAI struct {
	SliceType      int    // SST
	SliceDifferentiator string // SD (optional)
}

func main() {
	// Example NSSAI entries
	nssaiList := []NSSAI{
		{SliceType: 1, SliceDifferentiator: "0x010203"}, // eMBB
		{SliceType: 2, SliceDifferentiator: "0x112233"}, // URLLC
		{SliceType: 3},                                // mMTC
	}

	// Simulate network slice allocation
	for _, slice := range nssaiList {
		if slice.SliceDifferentiator != "" {
			fmt.Printf("Allocating slice: SST=%d, SD=%s\n", slice.SliceType, slice.SliceDifferentiator)
		} else {
			fmt.Printf("Allocating slice: SST=%d (default SD)\n", slice.SliceType)
		}
	}
}

Output:
Allocating slice: SST=1, SD=0x010203
Allocating slice: SST=2, SD=0x112233
Allocating slice: SST=3 (default SD)


Applications of NSSAI in 5G
Healthcare: URLLC slices for remote surgeries.
Automotive: URLLC slices for autonomous vehicles.
Entertainment: eMBB slices for high-definition video streaming.
IoT: mMTC slices for smart cities or industrial IoT.
