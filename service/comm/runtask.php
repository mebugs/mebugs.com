<?php
include_once($_SERVER['DOCUMENT_ROOT'].'/service/comm/doConfig.php');
include_once($_SERVER['DOCUMENT_ROOT'].'/service/makefile/makeFileTask.php');
// 启动定询任务
function autoRun() {
  //读取任务状态和启动时间以及最后执行时间
  $runStatus = true;
  // 扫描间隔
  $time = 10;
  do {
    $taskStatus = getConfig("taskStatus","./comm/system.php","int");
    $taskTime = getConfig("taskTime","./comm/system.php","int");
    $taskLastRun = getConfig("taskLastRun","./comm/system.php","string");
    $cdnUrl = getConfig("cdnUrl","./comm/system.php","string");
    $conns = [
      'qqUrl' => getConfig("qqUrl","./comm/system.php","string"),
      'githubUrl' => getConfig("githubUrl","./comm/system.php","string"),
      'giteeUrl' => getConfig("giteeUrl","./comm/system.php","string"),
      'emailUrl' => getConfig("emailUrl","./comm/system.php","string")
    ];
    if($taskStatus == 0) {
      $runStatus = false;
    } else {
      // 判断今天是不是执行过了，截取前10位 XXXX-XX-XX XX:XX:XX
      $lastRunDay = substr($taskLastRun,0,10);
      // 当前时间
      ini_set('date.timezone', 'Asia/Shanghai');
      $time_str = date('Y-m-d H:i:s', time());
      $today = substr($time_str,0,10);
      if($lastRunDay != $today) {
        $hour = substr($time_str,11,2);
        if($hour >= $taskTime) { // 执行今天的任务
          makeFileTask($cdnUrl,$today,$conns);// 处理全局文件生成
          setConfig("taskLastRun",$time_str,"./comm/system.php","string");// 更新执行时间
        }
      }
    }
    sleep($time);
    // 测试代码(windows下CGI不生效)
    $runStatus = false;
  } while ($runStatus);
}

?>