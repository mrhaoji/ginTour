from sys import exit

import time

def gold_room():
    print("Loading...")
    time.sleep(3)
    print("这个房间充满了RMB。 你带多少钱？(输入数字哦亲~)")
    choice = input('>')
    how_much = int(choice)

    if how_much <= 10000000000000:
        print("哦！你一点儿也不贪，你赢了~")
        jieshu_room()
    else:
        dead("你的贪婪害死了你！")
		
def sheep_room():
    print("Loading...")
    time.sleep(3)
    print("其实不需要加载，这个运行还是很快的。")
    time.sleep(2)
    print('这样会显得很高级。。。')
    time.sleep(2)
    print('继续继续。。。')
    time.sleep(1)
    print("这儿有一只羊。")
    time.sleep(2)
    print("这只羊旁边还有一捆草。")
    time.sleep(2)
    print("这只羊挡住了那扇门。")
    time.sleep(2)
    print("咋办？\n(拿走草/做鬼脸/踢羊屁股/打开门)")
    sheep_moved = False
	
    while True:
        choice = input(">")
		
        if choice == "拿走草":
            time.sleep(1)
            print("这只羊理撇了你一眼，它不喜欢吃草。并一脑壳撞了你一下，又回到了原来的位置。")
        elif choice == "踢羊屁股" and not sheep_moved:
            time.sleep(1)
            print("这只羊的屁股挨了一下，咩了一声跑开了。你可以进那个门了。")
            sheep_moved = True
        elif choice == "做鬼脸" and not sheep_moved:
            time.sleep(1)
            print("试问，这有什么用呢？是你傻还是羊傻？")
        elif choice == "打开门" and not sheep_moved:
            time.sleep(1)
            print("傻X，没看见那只羊啊。\n你怎么进的去？")
        elif choice =="打开门" and sheep_moved:
            time.sleep(1)
            print("在那只羊惨兮兮的注视下你走了进去。")
            gold_room()
        else:
            time.sleep(1)
            print("你他妈在说什么？？？\n看不到选项吗！！！")
			
def jieshu_room():
    time.sleep(1)
    print("结束咯咯咯~")
    time.sleep(1)
    print("欢迎使用。")
    time.sleep(1)
    print("请你问我打分~\n(1~5)")
    choice3 = input('>')
    fenshu = int(choice3)
    time.sleep(2)
    if fenshu != 5:
        dead('没有五分你就去shi吧。╭(╯^╰)╮')
    else:
        print('我会继续加油的。^_^')
        print('拜拜┏(＾0＾)┛')
        exit(0)

def wenda_room():
    print("Loading...")
    time.sleep(3)
    dead("你死了，这条路我还没想好怎么编~")


def dead(why):
    time.sleep(1)
    print(why,"你死掉了。Good job!!!")
    time.sleep(5)
    print("如果你愿意复活的话也可以哦~\n不过暂时只可以从头开始，从死的地方复活我还不会。\n(yes/no)")
    choice2 = input('>')
    time.sleep(2)
    if choice2 == 'yes':
        print("那重新开始了哦")
        time.sleep(3)
        start()
    else :
        print("那你是真的死了。")
        exit(0)
    
def start():
    print("你在一个黑暗的房间。")
    time.sleep(3)
    print("左右两边都有一条走廊。")
    time.sleep(3)
    print("你怎样选择。\n(左/右）")
	
    choice1 = input('>')
    if choice1 == '左':
        sheep_room()
    elif choice1 == '右':
        wenda_room()
    else:
        print("你是看不懂怎样输入，还是傻？")
        start()


print('这是王德智同学做的文字小游戏。\n《play4.0》')
time.sleep(3)
print('它很简单，但还是有很多不足，还有待完善。')
time.sleep(3)
print("玩的时候请对你的电脑好点，你可能想砸掉它。")
time.sleep(3)
print('那我们开始吧~~~~')
start()