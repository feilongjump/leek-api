<?php

namespace App\Policies;

use App\Models\User;
use App\Models\Project;
use App\Models\ProjectColumn;

class ProjectColumnPolicy extends Policy
{
    /**
     * Determine whether the user can create models.
     *
     * @param  \App\Models\User  $user
     * @param  \App\Models\ProjectColumn  $project
     * @return \Illuminate\Auth\Access\Response|bool
     */
    public function create(User $user)
    {
        return $user->tokenCan('project.column:create');
    }
}
