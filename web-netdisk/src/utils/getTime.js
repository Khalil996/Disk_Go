// getTime 负责用户登入进来后，说 上午|中午|下午|晚上|半夜|凌晨|早上 好
export const getTime = ()=> {
  const hours = new Date().getHours()
  if (hours < 3)
    return '深夜'
  else if (hours < 6)
    return '凌晨'
  else if (hours < 9)
    return '早上'
  else if (hours < 12)
    return '上午'
  else if (hours < 14)
    return '中午'
  else if (hours < 18)
    return '下午'
  else
    return '晚上'
}
