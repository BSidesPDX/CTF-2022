package com.github.blahblahbloopster.plugins

import io.ktor.server.routing.*
import io.ktor.http.*
import io.ktor.server.application.*
import io.ktor.server.response.*
import io.ktor.server.request.*
import java.io.File

const val FLAG = "BSidesPDX{N0m9gG485z0=}"

fun Application.configureRouting() {
    val basePage = this::class.java.classLoader.getResourceAsStream("static/index.html")!!.readAllBytes().decodeToString()
    // jank, but it works
    val index = basePage.replace("<!--.+-->".toRegex(), "")
    val alertBase = basePage.replace("<!--", "").replace("-->", "")
    val success = alertBase.replace("CONTENT", "You got the flag! <pre>$FLAG</pre>")
    val fail = alertBase.replace("CONTENT", "Not the flag!").replace("success", "danger")

    routing {
        get("/") {
            call.respondText(index, ContentType.Text.Html)
        }

        post("/flag") {
            val values = call.receiveParameters()
            val submission = values["flag"] ?: run { call.respond(400); return@post }
            if (insecureStringComp(submission, FLAG)) {
                call.respondText(success, ContentType.Text.Html)
            } else {
                call.respondText(fail, ContentType.Text.Html)
            }
        }
    }
}

fun insecureStringComp(a: String, b: String): Boolean {
    if (a.length != b.length) return false
    for ((c, d) in a.zip(b)) {
        Thread.sleep(100L)
        if (c != d) return false
    }
    return true
}
