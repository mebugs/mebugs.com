var doTeamIndex = setTeam;
var doTeams = ["tm_new","tm_up","tm_hot","tm_talk","tm_best"];
// 初始化下标
ready(initTeam)
function initTeam() {
  bodyNode.classList.add(doTeams[setTeam]);
}
// PC菜单浮动
function doTeam(index) {
  if(index != doTeamIndex) {
    bodyNode.classList.remove(doTeams[doTeamIndex]);
    doTeamIndex = index;
    bodyNode.classList.add(doTeams[doTeamIndex]);
  }
}
// PC菜单恢复
function doTeamOut() {
  if(doTeamIndex!=setTeam) {
    bodyNode.classList.remove(doTeams[doTeamIndex]);
    doTeamIndex = setTeam;
    bodyNode.classList.add(doTeams[doTeamIndex]);
  }
}