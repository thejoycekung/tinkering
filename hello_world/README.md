# A Guestbook

This is built using Python/Flask, with the support of a SQL DB and formatted (gently, truly), with Bootstrap. 

The underlying infrastructure:

1. A `Dockerfile` to create a container based on the available app files, which can then be pushed to Google Container Registry (GCR). 
2. Using Terraform (taken from Artemmkin's [Infrastructure as Code Tutorial](https://github.com/Artemmkin/infrastructure-as-code-tutorial)) you can then configure a Kubernetes cluster using GKE.
3. The `deployment.yaml` specifies the `LoadBalancer` service (to access the app outside of the cluster), a `PostgreSQL` pod+service with a `PersistentVolume` as the database, and of course a pod to run the app itself using our created image.

Post-mortem and walkthrough of the steps I took coming later!
