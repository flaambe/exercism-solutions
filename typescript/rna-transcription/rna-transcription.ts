export const DnaRnaMap: Record<string, string> = {
  G: 'C',
  C: 'G',
  T: 'A',
  A: 'U',
}

class Transcriptor {
  toRna(dna: string): string {
    let rna = ''

    for (const el of dna) {
      const nucleotide = DnaRnaMap[el]
      if (!nucleotide) throw Error('Invalid input DNA.')
      rna += nucleotide
    }

    return rna
  }
}

export default Transcriptor
