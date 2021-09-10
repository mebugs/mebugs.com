function AllTag() {
  return axios({
    method: 'post',
    url: '/service/tag.php',
    data: {api:'AllTag'}
  })
}