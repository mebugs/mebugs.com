<?php
//include(__DIR__.'/connect.php'); 
//校验TOKEN - 通过快速Token鉴权，暂不通过数据库
include($_SERVER['DOCUMENT_ROOT'].'/admin/comm/doConfig.php'); 
$servToken = getConfig("adminToken");
if($servToken != $token) {
	echo json_encode(['code' => 50008,'msg' => 'TOKEN非法' ,'data' => null]);
	exit;
}
?>