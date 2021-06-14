// install keda
test/registry.sh \
helm repo add kedacore https://kedacore.github.io/charts \
helm repo update \
kubectl create namespace keda \
helm install keda kedacore/keda --namespace keda 
// install fission
kubectl create namespace fission 
git clone --single-branch --branch master https://github.com/fission/fission.git 
cd fission/charts/fission-core 
helm dependency update . 
helm install fission-app . --set mqt_keda.enabled=true --set mqt_keda.connector_images.nats_steaming.image=localhost:5000/nats-steaming --set mqt_keda.connector_images.nats_steaming.tag=latest --set mqt_keda.connector_images.awssqs.image=localhost:5000/aws-sqs-test-connector --set mqt_keda.connector_images.awssqs.tag=latest --set mqt_keda.connector_images.kafka.image='localhost:5000/kafka-connector' --set mqt_keda.connector_images.kafka.tag=latest --set mqt_keda.connector_images.rabbitmq.image=localhost:5000/rabbit-keda --set mqt_keda.connector_images.rabbitmq.tag=latest -n fission
curl -Lo fission https://github.com/fission/fission/releases/download/1.11.1/fission-cli-linux && chmod +x fission && sudo mv fission /usr/local/bin/
