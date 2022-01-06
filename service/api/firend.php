<?php
// 读取站点信息
function CheckFirend($furl) {
  $data= readHtml($furl);
  if(!$data) {
    return [false,'站点数据读取失败'];
  }
  $meta=[];
  #title
  preg_match('/<title>([\w\W]*?)<\/title>/si', $data, $matches);
  if (!empty($matches[1])) {
    $meta['title'] = $matches[1];
  }
  #Description
  preg_match('/<meta\s+name="description"\s+content="([\w\W]*?)"/si', $data, $matches);
  if (empty($matches[1])) {
    preg_match("/<meta\s+name='description'\s+content='([\w\W]*?)'/si", $data, $matches);
  }
  if (empty($matches[1])) {
    preg_match('/<meta\s+content="([\w\W]*?)"\s+name="description"/si', $data, $matches);
  }
  if (empty($matches[1])) {
    preg_match('/<meta\s+http-equiv="description"\s+content="([\w\W]*?)"/si', $data, $matches);
  }
  if (!empty($matches[1])) {
    $meta['description'] = $matches[1];
  }
  #Icon
  preg_match('/<link\s+rel="icon"\s+href="([\w\W]*?)"/si', $data, $matches);
  if (empty($matches[1])) {
    preg_match("/<link\s+rel='icon'\s+href='([\w\W]*?)'/si", $data, $matches);
  }
  if (empty($matches[1])) {
    preg_match('/<link\s+href="([\w\W]*?)"\s+rel="icon"/si', $data, $matches);
  }
  if (empty($matches[1])) {
    preg_match("/<link\s+href='([\w\W]*?)'\s+rel='icon'/si", $data, $matches);
  }
  if (empty($matches[1])) {
    preg_match('/<link\s+rel="shortcut icon"\s+href="([\w\W]*?)"/si', $data, $matches);
  }
  if (empty($matches[1])) {
    preg_match("/<link\s+rel='shortcut icon'\s+href='([\w\W]*?)'/si", $data, $matches);
  }
  if (empty($matches[1])) {
    preg_match('/<link\s+href="([\w\W]*?)"\s+rel="shortcut icon"/si', $data, $matches);
  }
  if (empty($matches[1])) {
    preg_match("/<link\s+href='([\w\W]*?)'\s+rel='shortcut icon'/si", $data, $matches);
  }
  if (!empty($matches[1])) {
    $meta['icon'] = $matches[1];
  }
  return [true,$meta];
}


function readHtml($url) {
  $ch = curl_init();
  curl_setopt($ch, CURLOPT_URL, $url);
  curl_setopt($ch, CURLOPT_FOLLOWLOCATION, 1);
  curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
  curl_setopt($ch, CURLOPT_REFERER, $url);
  curl_setopt($ch, CURLOPT_USERAGENT, 'Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; GTB6.6; .NET CLR 2.0.50727; .NET CLR 3.0.04506.30; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729)');
  curl_setopt($ch, CURLOPT_HEADER, 0);
  curl_setopt($ch, CURLOPT_CONNECTTIMEOUT, 60);
  curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, 2);   // ssl 访问核心参数
  curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, FALSE); // ssl 访问核心参数
  $code = curl_exec($ch);
  curl_close($ch);
  if (empty($code) or stristr($code, '<title>403 Forbidden</title>') or stristr($code, "<hr><center>nginx") or stristr($code, "<title>Error</title>") or stristr($code, "<title>无法找到该页</title>")) {
      return false;
  } else {
      return $code;
  }    
}

// 提交友链
function AddFirend($conn,$body) {
  // 截取字符
  $fname = mb_substr($body -> fname,0,128,'UTF-8');
  $fdesc = mb_substr($body -> fdesc,0,128,'UTF-8');
  $furl = $body -> furl;
  $ficon = $body -> ficon;
  $findex = $body -> findex;
  $sql = "INSERT INTO `friend`(`furl`, `fname`, `fdesc`, `ficon`, `status`, `findex`) VALUES ('$furl', '$fname', '$fdesc', '$ficon', 0, '$findex')";
  $query = mysqli_query($conn,$sql);
  if(!$query) {
    return [false,'友链提交失败'];
  }
  return [true,'友链提交成功'];
}

// 获取友链管理员分页（级联文章）
function GetFirendsPageManage($conn,$q) {
  $ctSql = "SELECT count(id) FROM `friend` f";
  $qrSql = "SELECT * FROM `friend` f";
  $where = "";
  $ctSql = $ctSql . $where;
  $size = $q['size'];
  $qrSql = $qrSql . $where . " ORDER BY f.`status`,f.`id` DESC LIMIT ".($q['page']-1)*$size.",".$size;
  $pageData = [];
  $numAry = mysqli_fetch_array(mysqli_query($conn,$ctSql));
  $total = $numAry[0];
  $pageData['pages'] = ceil($total/$size);
  $query = mysqli_query($conn,$qrSql);
  $pageData['list'] = mysqli_fetch_all($query,MYSQLI_ASSOC);
  return [true,$pageData];
}

// 更新友链
function ModFirendsManage($conn,$update) {
  $sql = "UPDATE `friend` SET `furl` = '".$update['furl']."', `fname` = '".$update['fname']."', `fdesc` = '".$update['fdesc']."', `ficon` = '".$update['ficon']."', `status` = '".$update['status']."', `reson` = '".$update['reson']."', `findex` = '".$update['findex']."' WHERE `id` = ".$update['id'];
  $query = mysqli_query($conn,$sql);
  if($query) {
    return [true,$update['id']];
  }
  return [false,'数据更新失败'];
}
?>