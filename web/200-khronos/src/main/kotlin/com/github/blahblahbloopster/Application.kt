package com.github.blahblahbloopster

import io.ktor.server.engine.*
import io.ktor.server.netty.*
import com.github.blahblahbloopster.plugins.*

fun main() {
    embeddedServer(Netty, port = 80, host = "0.0.0.0") {
        configureRouting()
    }.start(wait = true)
}
