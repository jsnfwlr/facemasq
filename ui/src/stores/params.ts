import { defineStore } from 'pinia'
import { format } from 'date-fns'
import { mande } from 'mande'
import cloneDeep from 'lodash.clonedeep'


const params = mande('/')

export interface ParamState {
    Categories: Array<Category>;
    Statuses: Array<Status>;
    DeviceTypes: Array<DeviceType>;
    InterfaceTypes: Array<InterfaceType>;
    Locations: Array<Location>;
    Maintainers: Array<Maintainer>;
    OperatingSystems: Array<OperatingSystem>;
    Architectures: Array<Architecture>;
    VLANs: Array<VLAN>;
    Users: Array<User>;
}

export interface Category {
    ID: number | null;
    Label: string;
    Icon: string;
    Notes: string | null;
    IsLocked: boolean;

}
export interface Status {
    ID: number | null;
    Label: string;
    Icon: string;
    Notes: string | null;
    IsLocked: boolean;
}
export interface DeviceType {
    ID: number | null;
    Label: string;
    Icon: string;
    Notes: string | null;
    IsLocked: boolean;
}
export interface InterfaceType {
    ID: number | null;
    Label: string;
    Icon: string;
    Notes: string | null;
    IsLocked: boolean;
}
export interface Location {
    ID: number | null;
    Label: string;
    IsCloud: number;
    Notes: string | null;
    IsLocked: boolean;
}
export interface Maintainer {
    ID: number | null;
    Label: string;
    IsInternal: number;
    Notes: string | null;
    IsLocked: boolean;
}
export interface OperatingSystem {
    ID: number | null;
    Vendor: string;
    Family: string;
    Name: string;
    Version: string;
    IsOpenSource: boolean;
    IsServer: boolean;
    Notes: string | null;
    IsLocked: boolean;
}
export interface Architecture {
    ID: number | null;
    Label: string;
    BitSpace: number;
    Notes: string | null;
    IsLocked: boolean;
}
export interface VLAN {
    ID: number | null;
    Label: string;
    IPv4Mask: string;
    IPv6Mask: string;
    Notes: string | null;
    IsLocked: boolean;
}
export interface User {
    ID: number | null;
    Label: string;
    Password: string;
    NewPassword: string | null;
    Access: string;
    Notes: string | null;
    IsLocked: boolean;
}


export const useParams = defineStore('params', {
    state: () => {
        return ({
            Categories: [],
            Statuses: [],
            DeviceTypes: [],
            InterfaceTypes: [],
            Locations: [],
            Maintainers: [],
            OperatingSystems: [],
            Architectures: [],
            VLANs: [],
            Users: [],

        } as ParamState)
    },
    actions: {
        getParams() {
            params.get<ParamState>('/params').then((response) => {
                Object.keys(response).forEach((key) => {
                    switch (key) {
                        case 'Categories':
                            this.Categories = response.Categories
                            break
                        case 'Statuses':
                            this.Statuses = response.Statuses
                            break
                        case 'DeviceTypes':
                            this.DeviceTypes = response.DeviceTypes
                            break
                        case 'InterfaceTypes':
                            this.InterfaceTypes = response.InterfaceTypes
                            break
                        case 'Maintainers':
                            this.Maintainers = response.Maintainers
                            break
                        case 'Locations':
                            this.Locations = response.Locations
                            break
                        case 'OperatingSystems':
                            this.OperatingSystems = response.OperatingSystems
                            break
                        case 'Architectures':
                            this.Architectures = response.Architectures
                            break
                        case 'VLANs':
                            this.VLANs = response.VLANs
                            break
                        case 'Users':
                            this.Users = response.Users
                            break
                    }
                })
            })
        },

        NewCategory() {
            this.Categories.splice(0, 0, { ID: null, Label: "", Icon: "", Notes: null, IsLocked: false })
        },
        SaveCategory(index: number) {
            const isNew = ((this.Categories[index].ID === null) || (this.Categories[index].ID === 0))
            params.post('/categories', this.Categories[index]).then((response: any) => {
                this.Categories[index].ID = response.ID
                if (isNew) {
                    const move = this.Categories.splice(index, 1)
                    move.forEach((item => {
                        this.Categories.push(item)
                    }))
                }
            })
        },
        DropCategory(index: number) {
            this.Categories.splice(index, 1)
        },
        DeleteCategory(index: number) {
            params.delete('/categories/' + this.Categories[index].ID).then((response: any) => {
                this.Categories.splice(index, 1)
            })
        },

        NewStatus() {
            this.Statuses.splice(0, 0, { ID: null, Label: "", Icon: "", Notes: null, IsLocked: false })
        },
        SaveStatus(index: number) {
            const isNew = ((this.Statuses[index].ID === null) || (this.Statuses[index].ID === 0))
            params.post('/statuses', this.Statuses[index]).then((response: any) => {
                this.Statuses[index].ID = response.ID
                if (isNew) {
                    const move = this.Statuses.splice(index, 1)
                    move.forEach((item => {
                        this.Statuses.push(item)
                    }))
                }
            })
        },
        DropStatus(index: number) {
            this.Statuses.splice(index, 1)
        },
        DeleteStatus(index: number) {
            params.delete('/statuses/' + this.Statuses[index].ID).then((response: any) => {
                this.Statuses.splice(index, 1)
            })
        },

        NewMaintainer() {
            this.Maintainers.splice(0, 0, { ID: null, Label: "", Notes: null, IsInternal: 0, IsLocked: false })
        },
        SaveMaintainer(index: number) {
            const isNew = ((this.Maintainers[index].ID === null) || (this.Maintainers[index].ID === 0))

            params.post('/maintainers', this.Maintainers[index]).then((response: any) => {
                this.Maintainers[index].ID = response.ID
                if (isNew) {
                    const move = this.Maintainers.splice(index, 1)
                    move.forEach((item => {
                        this.Maintainers.push(item)
                    }))
                }
            })
        },
        DropMaintainer(index: number) {
            this.Maintainers.splice(index, 1)
        },
        DeleteMaintainer(index: number) {
            params.delete('/maintainers/' + this.Maintainers[index].ID).then((response: any) => {
                this.Maintainers.splice(index, 1)
            })
        },

        NewLocation() {
            this.Locations.splice(0, 0, { ID: null, Label: "", IsCloud: 0, Notes: null, IsLocked: false })
        },
        SaveLocation(index: number) {
            const isNew = ((this.Locations[index].ID === null) || (this.Locations[index].ID === 0))

            params.post('/locations', this.Locations[index]).then((response: any) => {
                this.Locations[index].ID = response.ID
                if (isNew) {
                    const move = this.Locations.splice(index, 1)
                    move.forEach((item => {
                        this.Locations.push(item)
                    }))
                }
            })
        },
        DropLocation(index: number) {
            this.Locations.splice(index, 1)
        },
        DeleteLocation(index: number) {
            params.delete('/locations/' + this.Locations[index].ID).then((response: any) => {
                this.Locations.splice(index, 1)
            })
        },

        NewDeviceType() {
            this.DeviceTypes.splice(0, 0, { ID: null, Label: "", Icon: "", Notes: null, IsLocked: false })
        },
        SaveDeviceType(index: number) {
            const isNew = ((this.DeviceTypes[index].ID === null) || (this.DeviceTypes[index].ID === 0))


            params.post('/deviceTypes', this.DeviceTypes[index]).then((response: any) => {
                this.DeviceTypes[index].ID = response.ID
                if (isNew) {
                    const move = this.DeviceTypes.splice(index, 1)
                    move.forEach((item => {
                        this.DeviceTypes.push(item)
                    }))
                }
            })
        },
        DropDeviceType(index: number) {
            this.DeviceTypes.splice(index, 1)
        },
        DeleteDeviceType(index: number) {
            params.delete('/deviceTypes/' + this.DeviceTypes[index].ID).then((response: any) => {
                this.DeviceTypes.splice(index, 1)
            })
        },

        NewOperatingSystem() {
            this.OperatingSystems.splice(0, 0, { ID: null, Vendor: "", Family: "", Name: "", Version: "", IsOpenSource: false, IsServer: false, Notes: null, IsLocked: false })
        },
        SaveOperatingSystem(index: number) {
            const isNew = ((this.OperatingSystems[index].ID === null) || (this.OperatingSystems[index].ID === 0))
            params.post('/operatingSystems', this.OperatingSystems[index]).then((response: any) => {
                this.OperatingSystems[index].ID = response.ID
                if (isNew) {
                    const move = this.OperatingSystems.splice(index, 1)
                    move.forEach((item => {
                        this.OperatingSystems.push(item)
                    }))
                }
            })
        },
        DropOperatingSystem(index: number) {
            this.OperatingSystems.splice(index, 1)
        },
        DeleteOperatingSystem(index: number) {
            params.delete('/operatingSystems/' + this.OperatingSystems[index].ID).then((response: any) => {
                this.OperatingSystems.splice(index, 1)
            })
        },

        NewArchitecture() {
            this.Architectures.splice(0, 0, { ID: null, Label: "", BitSpace: 0, Notes: null, IsLocked: false })
        },
        SaveArchitecture(index: number) {
            const isNew = ((this.Architectures[index].ID === null) || (this.Architectures[index].ID === 0))


            params.post('/architectures', this.Architectures[index]).then((response: any) => {
                this.Architectures[index].ID = response.ID
                if (isNew) {
                    const move = this.Architectures.splice(index, 1)
                    move.forEach((item => {
                        this.Architectures.push(item)
                    }))
                }
            })
        },
        DropArchitecture(index: number) {
            this.Architectures.splice(index, 1)
        },
        DeleteArchitecture(index: number) {
            params.delete('/architectures/' + this.Architectures[index].ID).then((response: any) => {
                this.Architectures.splice(index, 1)
            })
        },

        NewVLAN() {
            this.VLANs.splice(0, 0, { ID: null, Label: "", IPv4Mask: "", IPv6Mask: "", Notes: null, IsLocked: false })
        },
        SaveVLAN(index: number) {
            const isNew = ((this.VLANs[index].ID === null) || (this.VLANs[index].ID === 0))


            params.post('/vlans', this.VLANs[index]).then((response: any) => {
                this.VLANs[index].ID = response.ID
                if (isNew) {
                    const move = this.VLANs.splice(index, 1)
                    move.forEach((item => {
                        this.VLANs.push(item)
                    }))
                }
            })
        },
        DropVLAN(index: number) {
            this.VLANs.splice(index, 1)
        },
        DeleteVLAN(index: number) {
            params.delete('/vlans/' + this.VLANs[index].ID).then((response: any) => {
                this.VLANs.splice(index, 1)
            })
        },

        NewUser() {
            this.Users.splice(0, 0, { ID: null, Label: "", NewPassword: null, Password: "", Access: "", Notes: null, IsLocked: false })
        },
        SaveUser(index: number) {
            const isNew = ((this.Users[index].ID === null) || (this.Users[index].ID === 0))


            params.post('/users', this.Users[index]).then((response: any) => {
                this.Users[index].ID = response.ID
                if (isNew) {
                    const move = this.Users.splice(index, 1)
                    move.forEach((item => {
                        this.Users.push(item)
                    }))
                }
            })
        },
        DropUser(index: number) {
            this.Users.splice(index, 1)
        },
        DeleteUser(index: number) {
            params.delete('/users/' + this.Users[index].ID).then((response: any) => {
                this.Users.splice(index, 1)
            })
        },

    }
})
