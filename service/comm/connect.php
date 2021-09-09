<?php
//因为数据库仅有MYSQL 因此不采用PDO
//简单结构不采用读取配置文件的形式，直接给值
//后期如果需要做成安装包形式可以通过doConfig设置或获取配置信息
//include_once($_SERVER['DOCUMENT_ROOT'].'/service/comm/doConfig.php'); 
$servername = "localhost";
$username = "root";
$password = "123456";
$dbname = "mebugs_blog";
// 创建连接
$conn = new mysqli($servername, $username, $password,$dbname);
// 检测连接
if ($conn->connect_error) {
    die("连接失败: " . $conn->connect_error);
} 
mysqli_query($conn,"set character set 'utf8'");//读库 
mysqli_query($conn,"set names 'utf8'");//写库 
?>