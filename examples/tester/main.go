package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/FlowingSPDG/kairos-go"
	"github.com/hashicorp/go-multierror"
)

var (
	ip       = os.Getenv("KAIROS_IP")
	port     = os.Getenv("KAIROS_PORT")
	user     = os.Getenv("KAIROS_USER")
	password = os.Getenv("KAIROS_PASSWORD")
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	log.Println("KAIROSのAPIテストを開始します")
	log.Printf("テスト対象のKAIROS IP:%s PORT:%s USER:%s PASSWORD:%s\n", ip, port, user, password)

	kc := kairos.NewKairosRestClient(ip, port, user, password)
	// testing interface...

	var errs *multierror.Error

	log.Println("Aux Methodsのテストを開始します")
	auxErrors := testAuxMethods(ctx, kc)
	if auxErrors != nil {
		errs = multierror.Append(errs, auxErrors)
	}
	log.Println("Aux Methodsのテストが完了しました。 エラー:", auxErrors)

	log.Println("Inputsのテストを開始します")
	inputErrors := testInputs(ctx, kc)
	if inputErrors != nil {
		errs = multierror.Append(errs, inputErrors)
	}
	log.Println("Inputsのテストが完了しました。 エラー:", inputErrors)

	log.Println("Macrosのテストを開始します")
	macroErrors := testMacros(ctx, kc)
	if macroErrors != nil {
		errs = multierror.Append(errs, macroErrors)
	}
	log.Println("Macrosのテストが完了しました。 エラー:", macroErrors)

	log.Println("Multiviewersのテストを開始します")
	multiviewerErrors := testMultiviewers(ctx, kc)
	if multiviewerErrors != nil {
		errs = multierror.Append(errs, multiviewerErrors)
	}
	log.Println("Multiviewersのテストが完了しました。 エラー:", multiviewerErrors)

	log.Println("Scenesのテストを開始します")
	sceneErrors := testScenes(ctx, kc)
	if sceneErrors != nil {
		errs = multierror.Append(errs, sceneErrors)
	}
	log.Println("Scenesのテストが完了しました。 エラー:", sceneErrors)

	if errs != nil {
		log.Printf("テストに失敗しました。 \nエラー総数: %d", errs.Len())
		for i, err := range errs.Errors {
			log.Printf("エラー[%d]: %s", i, err.Error())
		}
	} else {
		log.Println("全てのテストが成功しました。")
	}
}

func testAuxMethods(ctx context.Context, kc kairos.KairosRestClient) (result *multierror.Error) {
	// Get Auxs
	auxs, err := kc.GetAuxs(ctx)
	if err != nil {
		result = multierror.Append(result, err)
	}

	// Get Aux
	// TODO...
	_ = auxs
	return
}

func testInputs(ctx context.Context, kc kairos.KairosRestClient) (result *multierror.Error) {
	// Get Inputs
	inputs, err := kc.GetInputs(ctx)
	if err != nil {
		result = multierror.Append(result, err)
	}

	// Get Input
	for _, input := range inputs {
		// test by ID
		_, err := kc.GetInputByID(ctx, input.UUID)
		if err != nil {
			result = multierror.Append(result, err)
		}

		// test by number
		_, err = kc.GetInputByNumber(ctx, input.Index)
		if err != nil {
			result = multierror.Append(result)
		}
	}
	return
}

func testMacros(ctx context.Context, kc kairos.KairosRestClient) (result *multierror.Error) {
	// Get Macros
	macros, err := kc.GetMacros(ctx)
	if err != nil {
		result = multierror.Append(result, err)
	}

	// Get Macro
	for _, macro := range macros {
		_, err := kc.GetMacro(ctx, macro.UUID)
		if err != nil {
			result = multierror.Append(result, err)
		}
	}
	return
}

func testMultiviewers(ctx context.Context, kc kairos.KairosRestClient) (result *multierror.Error) {
	// Get Multiviewers
	multiviewers, err := kc.GetMultiviewers(ctx)
	if err != nil {
		result = multierror.Append(result, err)
	}

	// Get Multiviewer
	// TODO...
	_ = multiviewers
	return
}

func testScenes(ctx context.Context, kc kairos.KairosRestClient) (result *multierror.Error) {
	// Get Scenes
	scenes, err := kc.GetScenes(ctx)
	if err != nil {
		result = multierror.Append(result, err)
	}

	// Get Scene
	for _, scene := range scenes {
		_, err := kc.GetScene(ctx, scene.UUID)
		if err != nil {
			result = multierror.Append(result, err)
		}
	}
	return
}
