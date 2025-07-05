https://docs.google.com/document/d/1BN-dj4YY_EDEj79fe3n2X7xPoet1bYWawzdr_uxGf0E/edit?tab=t.zqccwuuguyd

Network analytics information

The PCF shall be able to collect directly network analytic information from the NWDAF. 
The NWDAF provides network data analytics (e.g. load level information on a network slice level) to PCF. 
The PCF shall be able to use those data in its policy decisions. 
The details are defined in clause 6.1.1.3. 


NWDAF (Network Data Analytics Function) in 5G collects and analyzes network data to provide insights for optimizing network performance and improving user experience.

6.1.1.3 - Policy decisions based on network analytics
Policy decisions based on network analytics allow PCF to perform policy decisions taking into account the analytics information provided by the NWDAF. 
The PCF subscribes/unsubscribes to Analytics information as defined in TS 23.288 [24].
The following Analytics IDs are relevant for Policy decisions: "Load level information", "Service Experience","Network Performance" and "Abnormal behaviour". 
The PCF may subscribe to notifications of network analytics related to "Load Level Information" using the Nnwdaf_AnalyticsSubscription_Subscribe service operation including the Analytics ID "Load level information", the Analytics Filter "S-NSSAI and NSI ID" and the Analytics ReportingInformation set to a load level threshold value. 
The PCF is notified when the load level of the Network Slice Instance reaches the threshold, and then the PCF may verify if the RFSP index value needs to be modified for a SUPI for which an AM Policy Association is created; this is based on operator policies in the PCF, as defined in clause 6.1.2.1.
The PCF may subscribe to notifications of network analytics related to "Service Experience" using the Nnwdaf_AnalyticsSubscription_Subscribe service operation including the Analytics ID "Service Experience", the Target of Analytics Reporting "any UE" and the Analytics Filter including one or more "Application ID(s)". 
The PCF is notified on the Service Experience statistics or predictions including, for each Application Id, the list of SUPIs for which Service Experience is provided. 
In addition, both spatial and time validity may be provided as well as the
confidence of the prediction. 
The PCF may check the 5QI values assigned to the Application, the number of UEs
affected and may use this as input to calculate and update the authorized QoS for a service data flow template.
The NWDAF service to retrieve the service experience (i.e. the average observed Service MoS) is described in clause 6.4 of TS 23.288 [24].
The PCF may subscribe to notifications of network analytics related to "Network Performance" using the Nnwdaf_AnalyticsSubscription_Subscribe service operation including the Analytics ID "Network Performance", the Target of Analytics Reporting "Internal Group Id" and the Analytics Filter including the Area of Interest. 
The PCF is notified on the Network Performance statistics or predictions including the Area of Interest and both the gNB status information and the predicted number of UEs in the Area of Interest. In addition, the confidence of the prediction may be provided. 
The PCF may use this information as input to calculate the background data transfer policies that are negotiated with the ASP, as defined in clause 6.1.2.4.
The NWDAF services to retrieve "Network Performance" as described in clause 6.6 of TS 23.288 [24].
The PCF may subscribe to notifications of network analytics related to "Abnormal behaviour" using the Nnwdaf_AnalyticsSubscription_Subscribe service operation including the Analytics ID "Abnormal behaviour", the Target of Analytics Reporting "SUPI", "Internal Group Id" or "any UE" and the Analytics Filter including the expected
analytics type or the list of Exceptions IDs and per each Exception Id a possible threshold and other Analytics Filter Information if needed. 
The list of Exception IDs is specified in TS 23.288 [24]. 
The PCF may use "Unexpected UE location" as input to determine the Service Area Restrictions defined in clause 6.1.2.1, "Suspicion of DDoS attack" or "Too frequent Service Access" to request the SMF to terminate the PDU session as defined in clause 6.1.3.6, "Wrong destination address" to perform gating of a service data flow as defined in clause 6.1.3.6 and "Unexpected longlive/large rate flows" to perform QoS related policies such as gating or policing as defined in clause 6.2.1.1.
The NWDAF services to retrieve UE related analytics are described in clause 6.7 of TS 23.288 [24].
The PCF may also use the network analytics as input to its policy decision to apply operator defined actions for example for the UE context(s) or PDU session(s). 



