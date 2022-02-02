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

export default class UserAgent {
  private userAgent: UserAgentType;

  constructor(userAgent: string) {
    this.userAgent = JSON.parse(userAgent) as UserAgentType;
  }

  isMobile(): boolean {
    return this.userAgent.mobile && !this.userAgent.desktop;
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
