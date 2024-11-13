package main

import (
	"context"
	"dagger/boops/internal/dagger"
	"time"
)

type Boops struct{
	Ctr *dagger.Container
}

func New() *Boops {
	return &Boops{
		Ctr: dag.Container().From("alpine"),
	}
}

func (m *Boops) All(
    ctx context.Context,
    // +optional
    // +defaultPath=.
    source *dagger.Directory,
) []*dagger.Container {
    return []*dagger.Container{
        m.Build(ctx, source, false),
        m.Clippy(ctx, source),
        m.Typos(ctx, source),
    }
}

func (m *Boops) FirstOfAll(
    ctx context.Context,
    // +optional
    // +defaultPath=.
    source *dagger.Directory,
) *dagger.Container {
	return m.All(ctx, source)[0]
}

func (m *Boops) Core(
    ctx context.Context,
    // +optional
    // +defaultPath=.
    source *dagger.Directory,
) []*dagger.Container {
    return []*dagger.Container{
        dag.Container().From("alpine/git").WithExec([]string{"sh", "-c", "sleep 10"}),
        dag.Container().From("ubuntu").WithExec([]string{"sh", "-c", "sleep 10"}),
        dag.Container().From("nginx").WithExec([]string{"sh", "-c", "sleep 10"}),
    }
}

func (m *Boops) FirstOfCore(
    ctx context.Context,
    // +optional
    // +defaultPath=.
    source *dagger.Directory,
) *dagger.Container {
	return m.Core(ctx, source)[0]
}

func (m *Boops) Build(ctx context.Context, source *dagger.Directory, somebool bool) *dagger.Container {
	time.Sleep(10 * time.Second)
	return m.Ctr
}
func (m *Boops) Clippy(ctx context.Context, source *dagger.Directory) *dagger.Container {
	time.Sleep(10 * time.Second)
	return m.Ctr
}
func (m *Boops) Typos(ctx context.Context, source *dagger.Directory) *dagger.Container {
	time.Sleep(10 * time.Second)
	return m.Ctr
}
