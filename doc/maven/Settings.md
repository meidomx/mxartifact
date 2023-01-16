# Setting Sample

### Settings in settings.xml in ~/.m2 folder

```xml
<settings>
    <mirrors>
        <mirror>
            <id>self-mirror</id>
            <name>Self-hosted Mirror Repository</name>
            <url>http://localhost:9051/maven</url>
            <mirrorOf>central</mirrorOf>
        </mirror>
    </mirrors>

    <servers>
        <server>
            <id>deployserver</id>
            <username>admin</username>
            <password>admin123</password>
        </server>
    </servers>
</settings>
```

### Settings in project POM file

#### For deployment

```xml
<distributionManagement>
    <repository>
        <id>deployserver</id>
        <url>http://localhost:9051/maven-releases</url>
    </repository>
    <snapshotRepository>
        <id>deployserver</id>
        <url>http://localhost:9051/maven-snapshots</url>
    </snapshotRepository>
</distributionManagement>
```
