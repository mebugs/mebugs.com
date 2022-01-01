<?php 
//读写配置文件数据值获取。 
//默认根据字符串读取提取''中或""中的内容（$type不传值）
//如果有第三个参数时为int时按照数字int处理。 
//$path = $_SERVER['DOCUMENT_ROOT']."/service/comm/config.php";
function getConfig($key, $file="./comm/config.php", $type="string") 
{ 
  $res = ["",""];
	if ($type=="int") 
	{ 
		$str = file_get_contents($file); 
		$config = preg_match("/" . $key . "=(.*);/", $str, $res); 
    if($res==null) {
      $res = ["",0];
    }
		Return $res[1]; 
	} 
	else 
	{ 
		$str = file_get_contents($file); 
		$config = preg_match("/" . $key . "=\"(.*)\";/", $str, $res); 
		if($res==null) 
		{ 
      $config = preg_match("/" . $key . "='(.*)';/", $str, $res); 
		} 
    if($res==null) {
      $res =  ["",""];
    }
		Return $res[1]; 
	} 
} 
 
function setConfig($key, $value, $file="./comm/config.php", $type="string") 
{ 
	$str = file_get_contents($file); 
	$str2=""; 
	if($type=="int") 
	{ 
		$str2 = preg_replace("/" . $key . "=(.*);/", $key . "=" . $value . ";", $str); 
	} 
	else 
	{ 
		$str2 = preg_replace("/" . $key . "=(.*);/", $key . "=\"" . $value . "\";",$str); 
	} 
	file_put_contents($file, $str2); 
} 
?> 
