<?php

namespace App\Policies;

use App\Models\User;
use App\Models\ProjectColumnCard;

class ProjectColumnCardPolicy extends Policy
{
    /**
     * Determine whether the user can view the model.
     *
     * @param  \App\Models\User  $user
     * @param  \App\Models\ProjectColumnCard  $card
     * @return \Illuminate\Auth\Access\Response|bool
     */
    public function view(User $user, ProjectColumnCard $card)
    {
        return $card->column->project->user_id == $user->id;
    }

    /**
     * Determine whether the user can create models.
     *
     * @param  \App\Models\User  $user
     * @return \Illuminate\Auth\Access\Response|bool
     */
    public function create(User $user)
    {
        return $user->tokenCan('project.column.card:create');
    }

    /**
     * Determine whether the user can update the model.
     *
     * @param  \App\Models\User  $user
     * @param  \App\Models\ProjectColumnCard  $card
     * @return \Illuminate\Auth\Access\Response|bool
     */
    public function update(User $user, ProjectColumnCard $card)
    {
        return $card->column->project->user_id == $user->id && $user->tokenCan('project.column.card:update');
    }

    /**
     * Determine whether the user can delete the model.
     *
     * @param  \App\Models\User  $user
     * @param  \App\Models\ProjectColumnCard  $card
     * @return \Illuminate\Auth\Access\Response|bool
     */
    public function delete(User $user, ProjectColumnCard $card)
    {
        return $card->column->project->user_id == $user->id && $user->tokenCan('project.column.card:delete');
    }
}
