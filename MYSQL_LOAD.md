    - name: Run storageclass-s1
      shell: kubectl create -f /home/ubuntu/nfs-storageclass-s1.yaml
    - name: Wait for pod to be in running state
      kubernetes.core.k8s_info:
        kind: Pod
        name: my-pod-name
        namespace: my-namespace
      register: pod_info
      until: pod_info.resources[0].status.phase == "Running"
      retries: 30
      delay: 10
      delegate_to: ubuntu01