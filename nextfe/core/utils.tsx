
export function hashCode(input: string) {
  var hash = 0;
  for (var i = 0; i < input.length; i++) {
    var code = input.charCodeAt(i);
    hash = ((hash<<5)-hash)+code;
    hash = hash & hash; // Convert to 32bit integer
  }
  return hash + (2**31)
}

export function pickRandColor(colorNum: number, colors: number){
    if (colors < 1) colors = 1; // defaults to one color - avoid divide by zero
    return "hsl(" + (colorNum * (360 / colors) % 360) + ",100%,50%)";
}

