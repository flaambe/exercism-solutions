const Planet = {
  Earth: 1.0,
  Mercury: 0.2408467,
  Venus: 0.61519726,
  Mars: 1.8808158,
  Jupiter: 11.862615,
  Saturn: 29.447498,
  Uranus: 84.016846,
  Neptune: 164.79132,
} as const

const EarthPeriodInSeconds = 31_557_600

class SpaceAge {
  constructor(public readonly seconds: number) {}

  onEarth = (): number => this.convert(Planet.Earth)
  onMercury = (): number => this.convert(Planet.Mercury)
  onVenus = (): number => this.convert(Planet.Venus)
  onMars = (): number => this.convert(Planet.Mars)
  onJupiter = (): number => this.convert(Planet.Jupiter)
  onSaturn = (): number => this.convert(Planet.Saturn)
  onUranus = (): number => this.convert(Planet.Uranus)
  onNeptune = (): number => this.convert(Planet.Neptune)

  private convert(ratio: number): number {
    return Math.round((this.seconds / (EarthPeriodInSeconds * ratio)) * 100) / 100
  }
}

export default SpaceAge
