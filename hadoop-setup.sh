#!/bin/sh

set -e

sudo apt-get update
sudo apt-get install default-jdk
java -version

wget http://mirrors.tuna.tsinghua.edu.cn/apache/hadoop/common/hadoop-2.9.2/hadoop-2.9.2.tar.gz
tar -xzvf hadoop-2.9.2.tar.gz
sudo mv hadoop-2.9.2 /usr/local/hadoop

JAVAHOME = readlink -f /usr/bin/java | sed "s:bin/java::"
echo "Java Home : "$JAVAHOME

sed -i 's|${JAVA_HOME}|$(JAVAHOME)|g' /usr/local/hadoop/etc/hadoop/hadoop-env.sh

sudo tee /usr/local/hadoop/etc/hadoop/core-site.xml <<EOF
<configuration>
  <property>
    <name>fs.defaultFS</name>
    <value>hdfs://localhost:9000</value>
  </property>
</configuration>
EOF

sudo tee /usr/local/hadoop/etc/hadoop/hdfs-site.xml <<EOF
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
   <name>dfs.permissions.superusergroup</name>
   <value>hadoop</value>
  </property>
</configuration>
EOF

sudo mkdir -p /opt/hdfs/data /opt/hdfs/name

sudo /usr/local/hadoop/sbin/hadoop-daemon.sh start datanode
sudo /usr/local/hadoop/sbin/hadoop-daemon.sh start namenode
