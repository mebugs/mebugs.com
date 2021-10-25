function TagWork(data) {
  return axios({
    method: 'post',
    url: '/service/tag.php',
    data: data
  })
}

function ImgWork(data) {
  return axios({
    method: 'post',
    url: '/service/img.php',
    data: data
  })
}

function PostWork(data) {
  return axios({
    method: 'post',
    url: '/service/post.php',
    data: data
  })
}

function CategoryWork(data) {
  return axios({
    method: 'post',
    url: '/service/category.php',
    data: data
  })
}