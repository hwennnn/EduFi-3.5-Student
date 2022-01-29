import dateFormat from "dateformat";

export function formatDateStringFromMs(ms) {
    var dt = new Date(ms);

    return dateFormat(dt, "yyyy-mm-dd");
}