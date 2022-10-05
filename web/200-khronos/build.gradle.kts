val ktor_version: String by project
val kotlin_version: String by project
val logback_version: String by project

plugins {
    application
    kotlin("jvm") version "1.7.20"
    id("io.ktor.plugin") version "2.1.2"
}

group = "com.github.blahblahbloopster"
version = "0.0.1"
application {
    mainClass.set("com.github.blahblahbloopster.ApplicationKt")

    val isDevelopment: Boolean = project.ext.has("development")
    applicationDefaultJvmArgs = listOf("-Dio.ktor.development=$isDevelopment")
}

repositories {
    mavenCentral()
}

dependencies {
    implementation("io.ktor:ktor-server-core-jvm:$ktor_version")
    implementation("io.ktor:ktor-server-host-common-jvm:$ktor_version")
    implementation("io.ktor:ktor-server-netty-jvm:$ktor_version")
    implementation("ch.qos.logback:logback-classic:$logback_version")
    testImplementation("io.ktor:ktor-server-tests-jvm:$ktor_version")
    testImplementation("org.jetbrains.kotlin:kotlin-test-junit:$kotlin_version")
    testImplementation("io.ktor:ktor-client-core:$ktor_version")
    testImplementation("io.ktor:ktor-client-cio:$ktor_version")
}

// https://stackoverflow.com/a/68417096
tasks {
    jar {
        from(configurations.runtimeClasspath.get().map { if (it.isDirectory) it else zipTree(it) })

        duplicatesStrategy = DuplicatesStrategy.EXCLUDE

        manifest {
            attributes["Main-Class"] = "com.github.blahblahbloopster.ApplicationKt"
        }
    }
}
