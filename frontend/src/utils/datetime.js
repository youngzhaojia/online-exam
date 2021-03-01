export function formatDate(timeStamp) {
  if (!timeStamp) {
    return "";
  }
  timeStamp = timeStamp * 1000;
  const date = new Date(timeStamp * 1);
  const month = String(date.getMonth() + 1).padStart(2, "0");
  const dDate = String(date.getDate()).padStart(2, "0");
  return `${date.getFullYear()}-${month}-${dDate}`;
}

export function formatTime(timeStamp) {
  if (!timeStamp) {
    return "";
  }
  timeStamp = timeStamp * 1000;
  const date = new Date(timeStamp * 1);
  const hours = String(date.getHours()).padStart(2, "0");
  const minutes = String(date.getMinutes()).padStart(2, "0");
  const seconds = String(date.getSeconds()).padStart(2, "0");
  return `${hours}:${minutes}:${seconds}`;
}

export function formatDateTime(timeStamp) {
  return formatDate(timeStamp) + " " + formatTime(timeStamp);
}
