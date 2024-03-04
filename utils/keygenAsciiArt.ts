/**
 * Based on https://dev.to/krofdrakula/improving-security-by-drawing-identicons-for-ssh-keys-24mc
 */

interface Position {
  x: number;
  y: number;
}

const SYMBOLS = [
  ' ',
  '.',
  'o',
  '+',
  '=',
  '*',
  'B',
  'O',
  'X',
  '@',
  '%',
  '&',
  '#',
  '/',
  '^',
  'S',
  'E',
];
const WIDTH = 17;
const HEIGHT = 9;
const MOVES: Position[] = [
  {x: -1, y: -1}, // ↖
  {x: 1, y: -1}, // ↗
  {x: -1, y: 1}, // ↙
  {x: 1, y: 1}, // ↘
];

export class KeyGenAsciiArt {
  private commands: number[];
  private steps: number;
  constructor(commands: number[]) {
    this.commands = commands;
    this.steps = 0;
  }

  public *run(): Generator<string> {
    const commandLength = this.commands.length;
    while (this.steps < commandLength) {
      const world = simulate(this.commands, this.steps);

      yield draw(world, WIDTH, HEIGHT);

      this.steps += 1;
    }
  }
}

export const aa = (): [KeyGenAsciiArt, string] => {
  const key = randomHexString()
    .replace(/[^a-fA-F0-9]/g, '')
    .padStart(32, '0')
    .slice(0, 32);

  const commands = parseCommands(key);

  return [new KeyGenAsciiArt(commands), key];
};

const clamp = (min: number, max: number, x: number) => {
  return Math.max(min, Math.min(max, x));
};

const nextPosition = (position: Position, move: number) => {
  const delta = MOVES[move];
  return {
    x: clamp(0, WIDTH - 1, position.x + delta.x),
    y: clamp(0, HEIGHT - 1, position.y + delta.y),
  };
};

const splitByteIntoCommand = (byte: number): number[] => {
  return [byte & 3, (byte >>> 2) & 3, (byte >>> 4) & 3, (byte >>> 6) & 3];
};

const parseCommands = (hexString: string) => {
  const commands = [];
  // loop over all the characters in the hex string
  for (let i = 0; i < hexString.length; i += 2) {
    // take a pair of hex characters each time (one byte == 2 chars)
    const value = parseInt(hexString.slice(i, i + 2), 16);
    // split the byte into 4 double-bit numbers and append them to
    // the list of commands
    commands.push(...splitByteIntoCommand(value));
  }
  return commands;
};

const step = (
  world: number[],
  position: Position,
  command: number
): [number[], Position] => {
  // create a copy of the world state
  const newWorld = Array.from(world);
  // drop a coin in the current position
  newWorld[position.y * WIDTH + position.x] += 1;
  // return the new world state and the next position
  return [newWorld, nextPosition(position, command)];
};

const simulate = (commands: number[], steps = commands.length): number[] => {
  const start: Position = {x: 8, y: 4};
  let position = start;
  let world = Array(WIDTH * HEIGHT).fill(0);

  for (let i = 0; i < steps; i++) {
    [world, position] = step(world, position, commands[i]);
  }

  const end = position;
  world[start.y * WIDTH + start.x] = 15;
  world[end.y * WIDTH + end.x] = 16;

  return world;
};

const draw = (world: number[], width: number, height: number): string => {
  const drawing = world.map(cell => SYMBOLS[cell % SYMBOLS.length]).join('');

  const result = ['+' + '-'.repeat(width) + '+'];
  for (let i = 0; i < height; i++)
    result.push('|' + drawing.slice(i * width, (i + 1) * width) + '|');

  result.push('+' + '-'.repeat(width) + '+');
  return result.join('\n');
};

const randomHexString = (): string => {
  let key = '';
  for (let i = 0; i < 32; i++)
    key += Math.floor(Math.random() * 16).toString(16);
  return key;
};

export const defaultTable = (): string => {
  const world = Array(WIDTH * HEIGHT).fill(0);
  return draw(world, WIDTH, HEIGHT);
};
