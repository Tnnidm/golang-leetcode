import os

title = "5. Longest Palindromic Substring"
goname = "longestPalindrome"

if title != "" and goname != "":
    nameStr = title.replace(" ", "").replace(".", "_")

    quesNo = int(nameStr[0:nameStr.find("_")])
    parentFile = str(100*int((quesNo-1)/100)+1).zfill(4) + "_" + str(100*int((quesNo-1)/100)+100).zfill(4)

    nameStr = str(quesNo).zfill(4) + nameStr[nameStr.find("_"):]
    if not os.path.exists(os.path.join("/csh/go/src/golang-leetcode", parentFile)):
        os.makedirs(os.path.join("/csh/go/src/golang-leetcode", parentFile))
    
    if not os.path.exists(os.path.join("/csh/go/src/golang-leetcode", parentFile, nameStr)):
        os.makedirs(os.path.join("/csh/go/src/golang-leetcode", parentFile, nameStr))
    
    
    file = open(os.path.join("/csh/go/src/golang-leetcode", parentFile, nameStr, goname + ".go"), "w")
    file.write("package " + goname + "\n")
    file.close()

    file = open(os.path.join("/csh/go/src/golang-leetcode", parentFile, nameStr, goname + "_test.go"), "w")
    file.write("package " + goname + "\n")
    file.close()