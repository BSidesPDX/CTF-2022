#written by Aceroni
#aceroni.com
import asyncio
import os

import discord
from discord.ext import commands

TOKEN = os.getenv('DISCORD_TOKEN')
intents = discord.Intents.all()
intents.members = True
intents.presences = True

bot = commands.Bot(command_prefix="!", intents=intents)


class Ctf(commands.Cog):
    def __init__(self,b):
        self.bot = b
        self.player_states = {}
        self.GAME_STATES = {
            "UNINITIATED": {
                "question":"Hello, {name} would you like to play a game?",
                "answers":[],
                "correct_answer":"YES",
                'next_state':"QUESTION_1",
                "incorrect_response":"aww schucks let me know if you want to play",
                "correct_response":"awesome! let me know when you are ready for the next question by using the !ctf command"
            },
            "QUESTION_1":{
                "question":"Jung qnl jnf gur svefg OFvqrfCQK rirag?",
                "answers":["Sevqnl, Bpgbore 7, 2011","Zbaqnl, Whyl 4, 2011","Sevqnl, Abirzore 9, 2012","Fngheqnl, Frcgrzore 28, 2013"],
                "correct_answer":"FRIDAY, OCTOBER 7, 2011",
                "next_state":"QUESTION_2",
                "incorrect_response":"I am sorry that is incorrect, use !ctf to try again",
                "correct_response":"Good job! type !ctf to get the next question"
            },
            "QUESTION_2":{
                "question":"23-8-1-20 23-1-19 20-8-5 20-9-20-12-5 15-6 20-8-5 20-1-12-11 7-9-22-5-14 2-25 7-5-14-5 11-9-13 1-20 20-8-5 6-9-18-19-20 2-19-9-4-5-19-16-4-24 5-22-5-14-20?",
                "answers":["3-15-22-5-18-20 3-1-12-12-9-14-7: 19-5-3-18-5-20-19 15-6 19-15-3-9-1-12 5-14-7-9-14-5-5-18-9-14-7 18-5-22-5-1-12-5-4!","12-5-22-5-12 21-16: 8-15-23 19-5-3-21-18-9-20-25 9-19-14’20 12-9-11-5 16-12-1-25-9-14-7 1 22-9-4-5-15 7-1-13-5","23-8-25 9-14-6-15-19-5-3 9-19 8-5-12-16-9-14-7 9-20 6-1-9-12… 1-14-4 8-15-23 20-15 6-9-24 9-20","15-16-5-14-9-14-7 18-5-13-1-18-11-19"],
                "correct_answer":"WHY INFOSEC IS HELPING IT FAIL… AND HOW TO FIX IT",
                "next_state":"QUESTION_3",
                "incorrect_response":"I am sorry that is incorrect, use !ctf to try again",
                "correct_response":"Good job! type !ctf to get the next question"
            },
            "QUESTION_3":{
                "question":".-- .... . .-. . / -.. .. -.. / - .... . / ..--- ----- ..--- ----- / -... ... .. -.. . ... .--. -.. -..- / . ...- . -. - / - .- -.- . / .--. .-.. .- -.-. . ..--..",
                "answers":["... -- .. - .... / -- . -- --- .-. .. .- .-.. / ... - ..- -.. . -. - / ..- -. .. --- -.","... -- .. - .... / -- . -- --- .-. .. .- .-.. / ... - ..- -.. . -. - / ..- -. .. --- -.","--- .-. . --. --- -. / -.-. --- -. ...- . -. - .. --- -. / -.-. . -. - . .-.",".--- --- . / ..-. .. - --.. .----. ... / --. .- .-. .- --. ."],
                "correct_answer":"ONLINE",
                "next_state":"QUESTION_4",
                "incorrect_response":"I am sorry that is incorrect, use !ctf to try again",
                "correct_response":"Good job! type !ctf to get the next question"
            },
            "QUESTION_4": {
                "question":"9 44 666 0 444 7777 0 8 44 33 0 222 44 2 444 777 6 2 66 0 666 333 0 8 44 33 0 222 333 7 0 777 33 888 444 33 9 0 22 666 2 777 3 0 333 666 777 0 22 7777 444 3 33 7777 7 3 99 0 8 44 444 7777 0 999 33 2 777?",
                "answers":["8 666 7 44 33 777 0 8 444 6 9999 33 66","6 2 4 4 444 33 0 5 2 88 777 33 4 88 444","S6 2 777 444 666 66 0 6 2 777 7777 222 44 2 555 33 55","6 444 222 44 2 33 555 0 555 33 444 22 666 9 444 8 9999"],
                "correct_answer":"MICHAEL LEIBOWITZ",
                "next_state":"FINISHED",
                "incorrect_response":"I am sorry that is incorrect, use !ctf to try again",
                "correct_response":"Good job! type !ctf to get  your prize"

            },
            "FINISHED":{
                "flag":"BSidesPDX{s0m3t1m3s_4_C7F_f33ls_l1k3_4_7r1v14l_pur5u17}"
            }
         }

    async def run_quiz(self,ctx,state):
        if state == "FINISHED":
            await ctx.send(self.GAME_STATES[state]["flag"])
            return
        channel = ctx.channel
        await ctx.send(self.GAME_STATES[state]["question"].format(name = ctx.author.name) + "\n".join(self.GAME_STATES[state]["answers"]))
        def check(m):
            return m.channel == channel
        msg = await self.bot.wait_for('message', check = check)

        if msg.content.upper() == self.GAME_STATES[state]["correct_answer"]:
            self.player_states[ctx.author.id] = self.GAME_STATES[state]["next_state"]
            await ctx.channel.send(self.GAME_STATES[state]["correct_response"])
        else:
            await ctx.channel.send(self.GAME_STATES[state]["incorrect_response"])

    @commands.Cog.listener()
    async def on_message(self,message):
        if "FLAG" in message.content.upper():
            await message.channel.send("Cmon, you didn't think it would be that easy did you?")

    @commands.command(name = "ctf")
    async def cmd_ctf(self,ctx):
        """
        Starts the quiz to receive the flag for BSidesPDX PDX CTF . Must be used in a direct message with the 0xBill the bot.
        """
        if ctx.author.bot == True:
            return
        if not isinstance(ctx.channel,discord.DMChannel):
            return
        if ctx.author.id not in self.player_states:
            self.player_states[ctx.author.id] = "UNINITIATED"
        await self.run_quiz(ctx,self.player_states[ctx.author.id])


@bot.event
async def on_ready():
    print(f'{bot.user.name} has connected to Discord!')

async def setup(bot):
    await bot.add_cog(Ctf(bot))

if __name__ == "__main__":
    asyncio.run(setup(bot))
    bot.run(TOKEN)
