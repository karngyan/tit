import { Notification } from "~/constants"

const state = {
  notification: {
    show: false,
    type: Notification.Info, // success/danger/info
    title: '',
    message: '',
  },
  tweets: {},
  tconfigs: {}
}

export default function () {
  return state;
};
