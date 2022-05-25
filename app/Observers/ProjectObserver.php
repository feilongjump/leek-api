<?php

namespace App\Observers;

use App\Models\Project;
use Illuminate\Support\Arr;

class ProjectObserver
{
    public function creating(Project $project)
    {
        $project->user_id = auth()->id() ?? 1;
    }

    public function created(Project $project)
    {
        $this->saveContent($project);
    }

    public function saving(Project $project)
    {
        if ($project->isClean()) $project[$project::UPDATED_AT] = now();
    }

    public function saved(Project $project)
    {
        $this->saveContent($project);
    }

    public function deleted(Project $project)
    {
        $project->columns()->delete();
    }

    private function saveContent(Project $project)
    {
        $type = request('type', 'markdown');

        $data = Arr::only(request('content', []), $type);

        $project->content()->updateOrCreate(['contentable_id' => $project->id], $data);

        $project->loadMissing('content');
    }
}
