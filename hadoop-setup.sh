#!/bin/sh

set -e

sudo apt-get update
sudo apt-get install -y default-jdk wget
java -version

wget https://www-us.apache.org/dist/hadoop/common/hadoop-2.9.2/hadoop-2.9.2.tar.gz
tar -xzf hadoop-2.9.2.tar.gz
sudo mv hadoop-2.9.2 /usr/local/hadoop

JAVAHOME=$(readlink -f /usr/bin/java | sed "s:bin/java::")
echo "Java Home : "$JAVAHOME

sed -i "s|\${JAVA_HOME}|${JAVAHOME}|g" /usr/local/hadoop/etc/hadoop/hadoop-env.sh

sudo tee /usr/local/hadoop/etc/hadoop/core-site.xml <<EOF
<?xml version="1.0"?>
<configuration>
  <property>
    <name>fs.defaultFS</name>
    <value>hdfs://localhost:9000</value>
  </property>
</configuration>
EOF

sudo tee /usr/local/hadoop/etc/hadoop/hdfs-site.xml <<EOF
<?xml version="1.0"?>
<configuration>
  <property>
    <name>dfs.namenode.name.dir</name>
    <value>/opt/hdfs/name</value>
  </property>
  <property>
    <name>dfs.datanode.data.dir</name>
    <value>/opt/hdfs/data</value>
  </property>
  <property>
    <name>dfs.replication</name>
    <value>1</value>
  </property>
  <property>
   <name>dfs.permissions.superusergroup</name>
   <value>hadoop</value>
  </property>
</configuration>
EOF

sudo tee /usr/local/hadoop/etc/hadoop/mapred-site.xml <<EOF
<?xml version="1.0"?>
<configuration>
    <property>
        <name>mapreduce.framework.name</name>
        <value>yarn</value>
    </property>
</configuration>
EOF

sudo tee /usr/local/hadoop/etc/hadoop/yarn-site.xml <<EOF
<?xml version="1.0"?>
<configuration>
    <property>
        <name>yarn.nodemanager.aux-services</name>
        <value>mapreduce_shuffle</value>
    </property>
    <property>
        <name>yarn.nodemanager.aux-services.mapreduce_shuffle.class</name>
        <value>org.apache.hadoop.mapred.ShuffleHandler</value>
    </property>
    <property>
        <name>yarn.resourcemanager.hostname</name>
        <value>hadoop-master</value>
    </property>
</configuration>
EOF

export HADOOP_HOME=/usr/local/hadoop
export PATH=$PATH:/usr/local/hadoop/bin:/usr/local/hadoop/sbin

echo "PATH : "$PATH

sudo groupadd hadoop
sudo adduser travis hadoop

ssh-keygen -t rsa -f ~/.ssh/id_rsa -P '' && cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys

sudo mkdir -p /etc/hadoop
sudo mkdir -p /opt/hdfs/data /opt/hdfs/name
sudo chown -R travis:hadoop /opt/hdfs
sudo -u travis /usr/local/hadoop/bin/hdfs namenode -format -nonInteractive

sudo cp /usr/local/hadoop/etc/hadoop/*.* /etc/hadoop

echo -e "\n"
sudo /usr/local/hadoop/sbin/start-dfs.sh

echo -e "\n"
