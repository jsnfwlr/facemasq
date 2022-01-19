import { defineStore } from 'pinia'
import { format } from 'date-fns'
import { mande } from 'mande'

const settings = mande('/settings/')
const setting = mande('/setting/')

export interface UserState {
    account: User;
    settings: Settings;
}

export interface User {
    ID: number;
    Username: string;
    AccessLevel: number;
}

export interface Settings {
    darkMode: boolean | null;
    lastDismissed: string;
}

const topics = [
    "details",
    "params",
    "devices",
    "users"
]

const actions = [
    "read",
    "write",
    "delete"
]

export interface UserSetting {
    Name: string;
    Value: string;
}

export const useUser = defineStore('user', {
    state: () => {
        return ({
            settings: {
                darkMode: null,
                lastDismissed: '2021-12-21 08:55:56',
            },
            account: {
                ID: 1,
                Username: 'jsnfwlr',
                AccessLevel: 1,
            }
        } as UserState)
    },
    actions: {
        hasAccess(topic: string, action: string) {
            return true
        },
        updateSetting(name: string, value: any) {
            switch (name) {
                case 'darkMode':
                    this.settings.darkMode = (value === "true")
                    break
                case 'lastDismissed':
                    this.settings.lastDismissed = value
                    break
            }

        },
        getSettings() {
            settings.get<{ Name: string, Value: string }[]>(this.account.ID).then((response) => {
                if (response !== null && response.length > 0) {
                    response.forEach((elem) => {
                        this.updateSetting(elem.Name, elem.Value)
                    })
                }

            })
        },
        saveSetting(name: string, value: any) {
            switch (typeof value) {
                case "boolean":
                    value = (value) ? "true" : "false"
                    break
                case "number":
                    value = value.toString()
                    break
            }

            setting.post<UserSetting>(this.account.ID, { Name: name, Value: value }).then((response) => {
                this.updateSetting(response.Name, response.Value)
            })
        },
        isDarkMode() {
            if (this.settings.darkMode === null) {
                return window.matchMedia("(prefers-color-scheme: dark)").matches ? true : false
            }
            return this.settings.darkMode ? true : false
        },
        theme() {
            return this.isDarkMode() ? "dark" : "light"
        },
        toggleDarkMode() {
            this.saveSetting('darkMode', !this.isDarkMode())
        },
        lastDismissedUnknown(date: Date) {
            this.saveSetting('lastDismissed', format(date, 'yyyy-MM-dd HH:mm:ss'))
        }
    }
})
