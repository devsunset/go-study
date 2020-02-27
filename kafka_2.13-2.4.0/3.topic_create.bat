set JAVA_HOME=C:\dev\Java\zulu8.31.0.1-jdk8.0.181-win_x64
bin\windows\kafka-topics.bat --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic test

