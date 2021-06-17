	kubectlCmdSim simulates the output of kubectl requests locally and can be used for simple testing. 
	The  simulated pod status data is stored in a local text file called "podData" in this format, 
	where each pod has a corresponding scale:

	pod1:0
	pod2:1
	pod3:0


	To read the status of the pod, call:  ./kubectlCmdSim status podName             ex: ./kubectlCmdSim status pod1

	To change the scale of the pod, call: ./kubectlCmdSim scale podName podScale     ex: ./kubectlCmdSim scale pod1 1

	To reset every pod status to zero, call: ./kubectlCmdSim reset
	
	To query the number of deployment type, call: ./kubectlCmdSim getDeploymentNumber connectorType   ex: ./kubeCtlCmdSim getDeploymentNumber connector
