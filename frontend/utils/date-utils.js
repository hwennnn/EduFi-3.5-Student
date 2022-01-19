import dateFormat from "dateformat";

export function formatDateStringFromMs(ms) {
    var dt = new Date(ms);
    console.log(dt.getFullYear())
    return dateFormat(dt, "yyyy-mm-dd");
}