package main



import (

    "log"

    "math/rand"

    "time"

    "github.com/hajimehoshi/ebiten/v2"

    "github.com/hajimehoshi/ebiten/v2/audio"

    "github.com/hajimehoshi/ebiten/v2/ebitenutil"

)



// Game representa o estado do jogo

type Game struct {

    cat         *ebiten.Image

    background  *ebiten.Image

    victims     []string

    pointAudio  *audio.Player

    dieAudio    *audio.Player

    score       int

    catPos      float64

    isGameOver  bool

    isStarted   bool

}



const (

    screenWidth  = 800

    screenHeight = 600

    gravity      = 0.5

)



var (

    victims = []string{"bird", "mice", "rabbit", "butterfly", "dog"}

)



func NewGame() *Game {

    g := &Game{

        victims:    victims,

        score:      0,

        isGameOver: false,

        isStarted:  false,

    }

   

    // Carregar imagens e sons aqui

    // catImg, _, err := ebitenutil.NewImageFromFile("cat.png")

    // if err != nil {

    //     log.Fatal(err)

    // }

    // g.cat = catImg

   

    return g

}



func (g *Game) Update() error {

    if !g.isStarted {

        if ebiten.IsKeyPressed(ebiten.KeySpace) {

            g.isStarted = true

        }

        return nil

    }



    if g.isGameOver {

        if ebiten.IsKeyPressed(ebiten.KeyR) {

            g.reset()

        }

        return nil

    }



    // Lógica de movimento do gato

    if ebiten.IsKeyPressed(ebiten.KeySpace) {

        g.catPos -= 5

    }

    g.catPos += gravity



    // Verificar colisões e atualizar pontuação aqui



    return nil

}



func (g *Game) Draw(screen *ebiten.Image) {

    // Desenhar background

    if g.background != nil {

        op := &ebiten.DrawImageOptions{}

        screen.DrawImage(g.background, op)

    }



    // Desenhar gato

    if g.cat != nil {

        op := &ebiten.DrawImageOptions{}

        op.GeoM.Translate(100, g.catPos)

        screen.DrawImage(g.cat, op)

    }



    // Desenhar pontuação e estados do jogo

    if !g.isStarted {

        ebitenutil.DebugPrint(screen, "Pressione ESPAÇO para começar")

    }

    if g.isGameOver {

        ebitenutil.DebugPrint(screen, "Game Over! Pressione R para reiniciar")

    }

    ebitenutil.DebugPrint(screen, "Pontos: "+string(g.score))

}



func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {

    return screenWidth, screenHeight

}



func (g *Game) reset() {

    g.score = 0

    g.catPos = screenHeight / 2

    g.isGameOver = false

}



func main() {

    ebiten.SetWindowSize(screenWidth, screenHeight)

    ebiten.SetWindowTitle("Cat Game")

   

    game := NewGame()

    if err := ebiten.RunGame(game); err != nil {

        log.Fatal(err)

    }

}