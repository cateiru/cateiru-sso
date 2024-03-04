export interface UserAgentType {
  name: string;
  version: string;
  os: string;
  os_version: string;
  device: string;
  mobile: boolean;
  tablet: boolean;
  desktop: boolean;
  bot: boolean;
  url: string;
  string: string;
}

export enum Device {
  Mobile,
  Desktop,
  Tablet,
  Unknown,
}

export default class UserAgent {
  private userAgent: UserAgentType;

  constructor(userAgent: string) {
    this.userAgent = JSON.parse(userAgent) as UserAgentType;
  }

  device(): Device {
    if (this.userAgent.mobile) {
      return Device.Mobile;
    } else if (this.userAgent.desktop) {
      return Device.Desktop;
    } else if (this.userAgent.tablet) {
      return Device.Tablet;
    }

    return Device.Unknown;
  }

  uniqName(): string {
    if (this.userAgent.device) {
      return this.userAgent.device;
    } else if (this.userAgent.name && this.userAgent.os) {
      return `${this.userAgent.name} (${this.userAgent.os})`;
    } else if (this.userAgent.name) {
      return this.userAgent.name;
    } else if (this.userAgent.os) {
      return this.userAgent.os;
    }

    return 'Unknown';
  }
}
