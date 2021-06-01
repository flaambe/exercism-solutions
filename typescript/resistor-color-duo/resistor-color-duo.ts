export const ColorArr = [
  'black',
  'brown',
  'red',
  'orange',
  'yellow',
  'green',
  'blue',
  'violet',
  'grey',
  'white',
] as const

export type Color = typeof ColorArr[number]

export class ResistorColor {
  constructor(private colors: Color[]) {
    if (colors.length < 2) {
      throw new Error('At least two colors need to be present')
    }
  }

  value = (): number => ColorArr.indexOf(this.colors[0]) * 10 + ColorArr.indexOf(this.colors[1])
}
